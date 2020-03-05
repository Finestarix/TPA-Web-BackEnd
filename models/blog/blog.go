package blog

import (
	"../../connection"
	"../core"
	"log"
	"time"
)

type Blog struct {
	ID        int
	UserID    int
	Title     string
	Content   string
	ViewCount int
	Image     string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&Blog{})

	log.Println("Initialize Blog Success")
}

func DropTableBlog() {
	database := connection.GetConnection()
	defer database.Close()

	database.DropTableIfExists(&Blog{})
	database.AutoMigrate(&Blog{})

	log.Println("Drop Database Success")
}

func GetAllBlog() []Blog {
	database := connection.GetConnection()
	defer database.Close()

	var blogs []Blog
	if core.ValidateAPIKey() == false {
		return blogs
	}
	database.Find(&blogs)

	return blogs
}

func GetBlogByID(id int) Blog {

	database := connection.GetConnection()
	defer database.Close()

	var blog Blog
	if core.ValidateAPIKey() == false {
		return blog
	}
	database.Where("id = ?", id).First(&blog)

	return blog
}

func GetRecommendedBlog(id int) []Blog{
	database := connection.GetConnection()
	defer database.Close()

	var blogs []Blog
	if core.ValidateAPIKey() == false {
		return blogs
	}
	database.Where("id != ?", id).Limit(4).Find(&blogs)

	return blogs
}

func InsertBlog(userID int, title string, content string, image string, category string) *Blog {
	database := connection.GetConnection()
	defer database.Close()

	if len(title) < 5 {
		return nil
	}

	newBlog := &Blog{
		UserID:    userID,
		Title:     title,
		Content:   content,
		ViewCount: 0,
		Image:     image,
		Category:  category,
	}
	database.Save(newBlog)

	log.Println("Insert New Blog Success")
	return newBlog
}

func UpdateBlog(id int, content string, image string, category string) Blog {
	database := connection.GetConnection()
	defer database.Close()

	var updateBlog Blog
	if core.ValidateAPIKey() == false {
		return updateBlog
	}
	database.
		Model(&updateBlog).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"Content":  content,
			"Image":    image,
			"Category": category,
		})

	log.Println("Update Blog Success")
	return updateBlog
}

func DeleteBlog(id int) *Blog {
	database := connection.GetConnection()
	defer database.Close()

	var blog Blog
	blog = GetBlogByID(id)

	if blog.ID == 0 {
		log.Println("Delete Blog Failed")
		return &blog
	}

	err := database.Delete(blog).Error

	if err != nil {
		panic("Error Delete Hotel !" + err.Error())
	}

	log.Println("Delete Blog Success")
	return &blog
}
