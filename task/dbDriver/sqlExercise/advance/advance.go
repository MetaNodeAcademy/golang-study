package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ////////////////模型定义///////////////////
// 用户
type User struct {
	gorm.Model
	Name    string
	Age     int
	PostNum int `gorm:"type:int;default:0;column:post_num"`
}

// 文章
type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserId        int
	CommentStatus string `gorm:"type:varchar(255);default:'open'"`
}

// 评论
type Comment struct {
	gorm.Model
	Content string `gorm:"type:varchar(255);not null"`
	PostId  int    `gorm:"column:post_id"`
}

func main() {
	dsn := "root:root@/gorm-test?charset=utf8&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	// fmt.Println("数据表创建完成")
	// db.Create(&User{Name: "小王", Age: 18})
	// db.Create(&User{Name: "小李", Age: 20})
	// db.Create(&Post{Title: "文章1", Content: "内容1", UserId: 1})
	// db.Create(&Post{Title: "文章2", Content: "内容2", UserId: 1})
	// db.Create(&Post{Title: "文章3", Content: "内容3", UserId: 2})
	// db.Create(&Comment{Content: "评论1", PostId: 1})
	// db.Create(&Comment{Content: "评论2", PostId: 1})
	// db.Create(&Comment{Content: "评论3", PostId: 2})

	//删除评论
	db.Delete(&Comment{Model: gorm.Model{ID: 19}})
	db.Delete(&Comment{Model: gorm.Model{ID: 20}})
	db.Delete(&Comment{Model: gorm.Model{ID: 21}})

	//查询用户1的所有文章及评论
	var posts []Post
	db.Preload("Comments").Where("user_id=?", 1).Find(&posts)
	for _, post := range posts {
		fmt.Printf("用户1对应的文章标题:%s,内容:%s\n", post.Title, post.Content)
	}
	//查询评论数量最多的文章信息,使用group by
	var result Result
	db.Table("posts").
		Select("posts.*, count(comments.id) as maxCount").
		Joins("left join comments on posts.id = comments.post_id").
		Group("posts.id").
		Order("maxCount desc").
		Limit(1).
		Scan(&result)
	fmt.Printf("评论最多的文章标题:%s,内容:%s,评论数量:%d\n", result.Post.Title, result.Post.Content, result.MaxCount)

}

// 定义查询结果结构体
type Result struct {
	Post     Post `gorm:"embedded"`        //`gorm:"embedded"` 这个映射一定要要，否则无法赋值  字段首字母也必须大写，否则无法导出
	MaxCount int  `gorm:"column:maxCount"` //`gorm:"column:maxCount"` 这个映射一定要要，否则无法赋值
}

/*
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
*/
func (p *Post) AfterSave(tx *gorm.DB) (err error) {
	fmt.Println("文章创建完成...")
	// 更新用户的 PostNum 字段
	result := tx.Model(&User{}).Where("id = ?", p.UserId).Update("post_num", gorm.Expr("post_num + 1"))
	if result.Error != nil {
		fmt.Printf("更新用户文章数量失败: %v\n", result.Error)
		return result.Error
	}
	fmt.Printf("成功更新用户%d的文章数量，影响行数: %d\n", p.UserId, result.RowsAffected)
	return
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	// fmt.Println("评论删除完成...")

	// 获取 PostId（处理可能为0的情况）
	postId := c.PostId
	// 如果 PostId 为 0，从数据库获取完整信息
	if postId == 0 {
		var comment Comment
		result := tx.Unscoped().Select("post_id").First(&comment, c.ID)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				fmt.Println("评论记录不存在")
				return nil
			}
			fmt.Printf("获取评论信息失败: %v\n", result.Error)
			return result.Error
		}
		postId = comment.PostId
	}

	//检查当前文章对应的的评论数量是否为0
	var count int64
	result := tx.Model(&Comment{}).Where("post_id = ?", postId).Count(&count)
	if result.Error != nil {
		fmt.Printf("查询文章对应的评论数量失败: %v\n", result.Error)
		return result.Error
	}
	if count == 0 {
		fmt.Println("当前文章对应的评论数量为0...")
		//更新文章的评论数量字段为0
		result := tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comment_status", "无评论")
		if result.Error != nil {
			fmt.Printf("更新文章的评论数量字段失败: %v\n", result.Error)
		}
	}
	return
}
