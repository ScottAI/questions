package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"questions/pkg/database"
	"questions/pkg/models"
)

func Tags(c *gin.Context) {
	h := gin.H{}
	id := c.Param("id")
	questions := []models.Question{}
	users := []models.User{}
	session := sessions.Default(c)
	user := session.Get("user")
	var userId int

	database.DB.Find(&users)

	for _, v := range users {
		if v.Username == user {
			userId = v.Id
		}
	}

	database.DB.Raw("SELECT questions.*, users.username, users.id, tags.name FROM questions inner join users on users.id = questions.user_id inner join taggings on questions.id = taggings.question_id inner join tags on tags.id = taggings.tag_id where tags.name=?", id).Preload("Tags").Preload("User").Order("questions.id desc").Find(&questions)

	h["questions"] = questions
	h["userId"] = userId
	h["user"] = user
	h["id"] = id
	c.HTML(http.StatusOK, "tag.tmpl.html", h)
}

func EditTag(c *gin.Context) {
	path := c.Query("next")
	id := c.Param("id")
	h := gin.H{}
	tags := models.Tag{}
	users := []models.User{}
	database.DB.Where("tags.id = ?", id).Find(&tags)
	session := sessions.Default(c)
	user := session.Get("user")
	var userId int

	database.DB.Find(&users)

	for _, v := range users {
		if v.Username == user {
			userId = v.Id
		}
	}

	h["id"] = id
	h["user"] = user
	h["userId"] = userId
	h["tags"] = tags
	h["path"] = path
	session.Save()
	c.HTML(http.StatusOK, "tagedit.tmpl.html", h)

}

func UpdateTag(c *gin.Context) {
	session := sessions.Default(c)
	host := c.Request.Header.Get("Host")
	next := c.PostForm("next")
	name := c.PostForm("name")
	id := c.PostForm("id")
	tagid, _ := strconv.Atoi(id)
	tags := []models.Tag{}
	database.DB.Model(&tags).Where("id = ?", tagid).Update("name", name)
	session.Save()
	c.Redirect(http.StatusFound, host+next)
}

func Categories(c *gin.Context) {
	type Result struct {
		Name  string
		Total int
	}
	results := []Result{}
	users := []models.User{}
	session := sessions.Default(c)
	user := session.Get("user")
	var userId int

	database.DB.Find(&users)

	for _, v := range users {
		if v.Username == user {
			userId = v.Id
		}
	}

	session.Save()
	database.DB.Raw("SELECT DISTINCT name, count(name) as total FROM tags GROUP BY name ORDER BY total DESC").Scan(&results)
	c.HTML(http.StatusOK, "categories.tmpl.html",
		gin.H{"results": results,
			"user":   user,
			"userId": userId,
		})
}
