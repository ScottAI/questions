package handlers

import (
	"math"
	"net/http"
	"questions/pkg/logger"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"questions/pkg/database"
	"questions/pkg/models"
)

func AllQuestions(c *gin.Context) {
	logger.Info("access /")
	h := gin.H{}
	questions := []models.Question{}
	users := []models.User{}
	session := sessions.Default(c)
	user := session.Get("user")
	p := c.Query("page")
	var count int
	database.DB.Find(&questions).Count(&count)

	var userId int

	database.DB.Find(&users)

	for _, v := range users {
		if v.Username == user {
			userId = v.Id
		}
	}

	if p == "" {
		p = "1"
	}
	page, _ := strconv.ParseInt(p, 10, 32)
	per := int64(3)
	totalPages := int(math.Ceil(float64(count) / float64(per)))
	offset := per * (page - 1)

	database.DB.Preload("Tags").
		Preload("User").
		Order("id desc").
		Limit(per).
		Offset(offset).
		Find(&questions)

	h["questions"] = questions
	h["user"] = user
	h["userId"] = userId
	h["totalPages"] = totalPages
	h["Flash"] = session.Flashes()
	c.HTML(http.StatusOK, "index.tmpl.html", h)
}

func UnsolvedQuestions(c *gin.Context) {
	h := gin.H{}
	questions := []models.Question{}
	users := []models.User{}
	session := sessions.Default(c)
	user := session.Get("user")
	p := c.Query("page")
	var count int
	database.DB.Where("accepted_answer = false").Find(&questions).Count(&count)

	var userId int

	database.DB.Find(&users)

	for _, v := range users {
		if v.Username == user {
			userId = v.Id
		}
	}

	if p == "" {
		p = "1"
	}
	page, _ := strconv.ParseInt(p, 10, 32)
	per := int64(3)
	totalPages := int(math.Ceil(float64(count) / float64(per)))
	offset := per * (page - 1)

	database.DB.Preload("Tags").
		Preload("User").
		Where("accepted_answer = false").
		Order("id desc").
		Limit(per).
		Offset(offset).
		Find(&questions)

	h["questions"] = questions
	h["user"] = user
	h["userId"] = userId
	h["totalPages"] = totalPages
	h["Flash"] = session.Flashes()
	c.HTML(http.StatusOK, "unsolved.tmpl.html", h)
}

func SolvedQuestions(c *gin.Context) {
	h := gin.H{}
	questions := []models.Question{}
	users := []models.User{}
	session := sessions.Default(c)
	user := session.Get("user")
	p := c.Query("page")
	var count int
	database.DB.Where("accepted_answer = true").Find(&questions).Count(&count)

	var userId int

	database.DB.Find(&users)

	for _, v := range users {
		if v.Username == user {
			userId = v.Id
		}
	}

	if p == "" {
		p = "1"
	}
	page, _ := strconv.ParseInt(p, 10, 32)
	per := int64(3)
	totalPages := int(math.Ceil(float64(count) / float64(per)))
	offset := per * (page - 1)

	database.DB.Preload("Tags").
		Preload("User").
		Where("accepted_answer = true").
		Order("id desc").
		Limit(per).
		Offset(offset).
		Find(&questions)

	h["questions"] = questions
	h["user"] = user
	h["userId"] = userId
	h["totalPages"] = totalPages
	h["Flash"] = session.Flashes()
	c.HTML(http.StatusOK, "solved.tmpl.html", h)
}

func MostViewedQuestions(c *gin.Context) {
	h := gin.H{}
	questions := []models.Question{}
	users := []models.User{}
	session := sessions.Default(c)
	user := session.Get("user")
	p := c.Query("page")
	var count int
	database.DB.Find(&questions).Count(&count)

	var userId int

	database.DB.Find(&users)

	for _, v := range users {
		if v.Username == user {
			userId = v.Id
		}
	}

	if p == "" {
		p = "1"
	}
	page, _ := strconv.ParseInt(p, 10, 32)
	per := int64(3)
	totalPages := int(math.Ceil(float64(count) / float64(per)))
	offset := per * (page - 1)

	database.DB.Preload("Tags").
		Preload("User").
		Order("templates desc").
		Limit(per).
		Offset(offset).
		Find(&questions)

	h["questions"] = questions
	h["user"] = user
	h["userId"] = userId
	h["totalPages"] = totalPages
	h["Flash"] = session.Flashes()
	c.HTML(http.StatusOK, "viewed.tmpl.html", h)
}

func OldestQuestions(c *gin.Context) {
	h := gin.H{}
	questions := []models.Question{}
	users := []models.User{}
	session := sessions.Default(c)
	user := session.Get("user")
	p := c.Query("page")
	var count int
	database.DB.Find(&questions).Count(&count)

	var userId int

	database.DB.Find(&users)

	for _, v := range users {
		if v.Username == user {
			userId = v.Id
		}
	}

	if p == "" {
		p = "1"
	}
	page, _ := strconv.ParseInt(p, 10, 32)
	per := int64(3)
	totalPages := int(math.Ceil(float64(count) / float64(per)))
	offset := per * (page - 1)

	database.DB.Preload("Tags").
		Preload("User").
		Order("id asc").
		Limit(per).
		Offset(offset).
		Find(&questions)

	h["questions"] = questions
	h["user"] = user
	h["userId"] = userId
	h["totalPages"] = totalPages
	h["Flash"] = session.Flashes()
	c.HTML(http.StatusOK, "oldest.tmpl.html", h)
}

func SearchQuestions(c *gin.Context) {
	h := gin.H{}
	questions := []models.Question{}
	users := []models.User{}
	session := sessions.Default(c)
	user := session.Get("user")
	q := c.Query("q")
	p := c.Query("page")
	var count int

	database.DB.Where("title LIKE ?", "%"+q+"%").Find(&questions).Count(&count)

	if p == "" {
		p = "1"
	}
	page, _ := strconv.ParseInt(p, 10, 32)
	per := int64(3)
	totalPages := int(math.Ceil(float64(count) / float64(per)))
	offset := per * (page - 1)

	var userId int

	database.DB.Find(&users)

	for _, v := range users {
		if v.Username == user {
			userId = v.Id
		}
	}

	database.DB.Preload("Tags").
		Preload("User").
		Where("title LIKE ?", "%"+q+"%").
		Order("id desc").
		Limit(per).
		Offset(offset).
		Find(&questions)

	h["questions"] = questions
	h["user"] = user
	h["userId"] = userId
	h["totalPages"] = totalPages
	h["q"] = q
	c.HTML(http.StatusOK, "search.tmpl.html", h)
}

func ShowQuestion(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	c.Writer.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	c.Writer.Header().Set("Expires", "0")                                         // Proxies.
	id := c.Param("id")
	question := models.Question{}
	tags := []models.Tag{}
	users := []models.User{}
	answers := []models.Answer{}
	session := sessions.Default(c)
	user := session.Get("user")
	var count int

	database.DB.Find(&users)

	var answerUserId int
	database.DB.Find(&users)
	for _, v := range users {
		if v.Username == user {
			answerUserId = v.Id
		}
	}

	database.DB.Exec("UPDATE questions SET templates = templates + 1 WHERE questions.id = ?", id)
	database.DB.Raw("SELECT tags.name FROM questions join taggings join tags on questions.id = taggings.question_id and tags.id = taggings.tag_id where questions.id=?", id).Find(&tags)
	database.DB.Preload("User").Where("questions.id=?", id).Find(&question)
	database.DB.Preload("User").
		Preload("Question").
		Where("answers.question_id=?", id).
		Group("answers.id").
		Order("id desc").
		Find(&answers).Count(&count)

	c.HTML(http.StatusOK, "show.tmpl.html",
		gin.H{"question": question,
			"tags":         tags,
			"user":         user,
			"answerUserId": answerUserId,
			"answers":      answers,
			"count":        count})
}

func CreateQuestion(c *gin.Context) {
	h := gin.H{}
	users := []models.User{}
	session := sessions.Default(c)
	user := session.Get("user")
	var questionUserID int
	database.DB.Find(&users)
	for _, v := range users {
		if v.Username == user {
			questionUserID = v.Id
		}
	}
	h["user"] = user
	h["questionUserID"] = questionUserID
	session.Save()
	c.HTML(http.StatusOK, "create.tmpl.html", h)
}

func SaveQuestion(c *gin.Context) {
	h := gin.H{}
	session := sessions.Default(c)
	u := c.PostForm("user")
	questionUserID, _ := strconv.Atoi(u)
	title := c.PostForm("title")
	body := c.PostForm("body")
	name := c.PostForm("name")
	session.Set("questionUserID", questionUserID)
	session.Set("name", name)
	session.Set("title", title)
	session.Set("body", body)
	h["title"] = title
	h["body"] = body
	h["name"] = name
	h["questionUserID"] = questionUserID

	if title == "" {
		h["a"] = "Required field can't be empty!"
		c.HTML(http.StatusFound, "create.tmpl.html", h)
		return
	}

	if name != "" {
		re := regexp.MustCompile("(.*?),")
		matched := re.Match([]byte(name))
		if matched == false {
			h["b"] = "Must be comma-separated!"
			c.HTML(http.StatusOK, "create.tmpl.html", h)
			return
		}
	}
	// for insert split tags string to db
	w := strings.TrimSpace(c.PostForm("name"))
	z := strings.Trim(w, ",")
	tagsList := []models.Tag{}
	for _, tag := range strings.Split(z, ",") {
		tagsList = append(tagsList, models.Tag{Name: tag})
	}

	questions := models.Question{
		UserID: questionUserID,
		Title:  title,
		Body:   body,
		Tags:   tagsList,
	}

	database.DB.Save(&questions)
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

func EditQuestion(c *gin.Context) {
	id := c.Param("id")
	session := sessions.Default(c)
	user := session.Get("user")
	question := models.Question{}
	tags := []models.Tag{}
	users := []models.User{}

	var questionUserID int
	database.DB.Find(&users)
	for _, v := range users {
		if v.Username == user {
			questionUserID = v.Id
		}
	}

	database.DB.Raw("SELECT tags.name, tags.id FROM questions join taggings join tags on questions.id = taggings.question_id and tags.id = taggings.tag_id where questions.id=?", id).Find(&tags)
	database.DB.Find(&question, id)

	c.HTML(http.StatusOK, "edit.tmpl.html",
		gin.H{"question": question,
			"tags":           tags,
			"user":           user,
			"questionUserID": questionUserID,
		})
}

func UpdateQuestion(c *gin.Context) {
	h := gin.H{}
	id := c.Param("id")
	session := sessions.Default(c)
	u := c.PostForm("user")
	userid, _ := strconv.Atoi(u)
	title := c.PostForm("title")
	body := c.PostForm("body")
	session.Set("title", title)
	session.Set("body", body)
	h["userid"] = userid
	h["title"] = title
	h["body"] = body

	if title == "" {
		h["a"] = "Required field can't be empty!"
		c.HTML(http.StatusFound, "edit.tmpl.html", h)
		return
	}

	question := models.Question{}
	database.DB.Find(&question, id)
	question.Title = title
	question.Body = body
	database.DB.Save(&question)
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

func DeleteQuestion(c *gin.Context) {
	ip := c.Request.Header.Get("Referer")
	id := c.Param("id")
	questions := []models.Question{}
	questionId, _ := strconv.Atoi(id)
	database.DB.Delete(&questions, questionId)
	c.Redirect(http.StatusFound, ip)
}

func QuestionLikes(c *gin.Context) {
	questions := []models.Question{}
	id := c.PostForm("id")
	questionId, _ := strconv.Atoi(id)
	database.DB.Model(&questions).
		Where("id = ?", questionId).
		UpdateColumn("likes", gorm.Expr("likes + ?", 1))

	t := strconv.Itoa(questionId)
	c.Redirect(http.StatusFound, "/show/"+t)
}

func Chat(c *gin.Context) {
	type Result struct {
		Username string
	}

	results := []Result{}
	database.DB.Raw("SELECT username FROM users WHERE is_logged_in = 1").Scan(&results)

	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	c.Writer.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	c.Writer.Header().Set("Expires", "0")                                         // Proxies.
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

	if user == nil {
		c.Redirect(http.StatusSeeOther, "/login")
	}

	c.HTML(http.StatusOK, "chat.tmpl.html", gin.H{
		"userId":  userId,
		"user":    user,
		"results": results,
	})
}
