package blog

import (
	models "../../models/blog"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllBlog(p graphql.ResolveParams) (i interface{}, err error) {
	blog := models.GetAllBlog()
	if len(blog) == 0 {
		return nil, errors.New("error: no data found")
	}

	return blog, nil
}

func InsertBlog(p graphql.ResolveParams) (i interface{}, err error) {
	userID := p.Args["userID"].(int)
	title := p.Args["title"].(string)
	content := p.Args["content"].(string)
	image := p.Args["image"].(string)
	category := p.Args["category"].(string)

	newBlog := models.InsertBlog(userID, title, content, image, category)

	return newBlog, nil
}

func UpdateBlog(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)
	content := p.Args["content"].(string)
	image := p.Args["image"].(string)
	category := p.Args["category"].(string)

	blog := models.UpdateBlog(id, content, image, category)

	return blog, nil
}

func DeleteBlog(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	blog := models.DeleteBlog(id)

	return blog, nil
}
