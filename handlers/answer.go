package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"questions/pkg/database"
	"questions/pkg/models"
)

func SaveAnswer(c *gin.Context) {
	ip := c.Request.Header.Get("Referer")
	session := sessions.Default(c)
	u := c.PostForm("user")
	i := c.PostForm("id")
	body := c.PostForm("body")

	answerUserId, _ := strconv.Atoi(u)
	questionUserId, _ := strconv.Atoi(i)
	answers := models.Answer{
		UserID:     answerUserId,
		QuestionID: questionUserId,
		Body:       body,
	}

	database.DB.Save(&answers)
	database.DB.Exec("UPDATE questions SET answer_count = answer_count + 1 WHERE questions.id = ?", questionUserId)

	session.Save()
	c.Redirect(http.StatusFound, ip)
}

func AcceptAnswer(c *gin.Context) {
	answer := []models.Answer{}
	question := []models.Question{}
	id := c.PostForm("qid")
	ans := c.PostForm("aid")
	answerId, _ := strconv.Atoi(ans)
	questionId, _ := strconv.Atoi(id)
	database.DB.Model(&answer).
		Where("answers.id = ?", answerId).
		UpdateColumn("is_accepted_answer", gorm.Expr("is_accepted_answer + ?", 1))

	database.DB.Model(&question).
		Where("questions.id = ?", questionId).
		UpdateColumn("accepted_answer", gorm.Expr("accepted_answer + ?", 1))

	t := strconv.Itoa(questionId)
	c.Redirect(http.StatusFound, "/show/"+t)
}

func EditAnswer(c *gin.Context) {
	id := c.Param("id")
	session := sessions.Default(c)
	user := session.Get("user")
	answer := models.Answer{}
	users := []models.User{}

	var answerUserID int
	database.DB.Find(&users)
	for _, v := range users {
		if v.Username == user {
			answerUserID = v.Id
		}
	}

	database.DB.Find(&answer, id)

	c.HTML(http.StatusOK, "answeredit.tmpl.html",
		gin.H{"answer": answer,
			"user":         user,
			"answerUserID": answerUserID,
		})
}

func UpdateAnswer(c *gin.Context) {
	id := c.Param("id")
	session := sessions.Default(c)
	body := c.PostForm("body")
	answer := models.Answer{}

	database.DB.Model(&answer).Where("id = ?", id).Update("body", body)
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

func AnswerLikes(c *gin.Context) {
	answer := []models.Answer{}
	id := c.PostForm("qid")
	ans := c.PostForm("aid")
	answerId, _ := strconv.Atoi(ans)
	questionId, _ := strconv.Atoi(id)
	database.DB.Model(&answer).
		Where("answers.id = ?", answerId).
		UpdateColumn("likes", gorm.Expr("likes + ?", 1))

	t := strconv.Itoa(questionId)
	c.Redirect(http.StatusFound, "/show/"+t)
}

func AnswerDisLikes(c *gin.Context) {
	answer := []models.Answer{}
	id := c.PostForm("qid")
	ans := c.PostForm("aid")
	answerId, _ := strconv.Atoi(ans)
	questionId, _ := strconv.Atoi(id)
	database.DB.Model(&answer).
		Where("answers.id = ?", answerId).
		UpdateColumn("dis_likes", gorm.Expr("dis_likes + ?", 1))

	t := strconv.Itoa(questionId)
	c.Redirect(http.StatusFound, "/show/"+t)
}

func AnswerDelete(c *gin.Context) {
	id := c.Param("id")
	ip := c.Request.Header.Get("Referer")
	answerId, _ := strconv.Atoi(id)
	answers := []models.Answer{}
	database.DB.Delete(&answers, answerId)
	c.Redirect(http.StatusFound, ip)
}
