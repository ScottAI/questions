package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"questions/handlers"
)

func RegisterRouters() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Logger())
	r.LoadHTMLGlob("templates/*.tmpl.html")

	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("qussession", store))

	//questions
	r.GET("/", handlers.AllQuestions)
	r.GET("/unsolved", handlers.UnsolvedQuestions)
	r.GET("/solved", handlers.SolvedQuestions)
	r.GET("/viewed", handlers.MostViewedQuestions)
	r.GET("/oldest", handlers.OldestQuestions)
	r.GET("/search", handlers.SearchQuestions)
	r.GET("/show/:id", handlers.ShowQuestion)
	r.GET("/create", handlers.CreateQuestion)
	r.GET("/edit/:id", handlers.EditQuestion)
	r.POST("/update/:id", handlers.UpdateQuestion)
	r.POST("/delete/:id", handlers.DeleteQuestion)
	r.POST("/savequestion", handlers.SaveQuestion)
	r.POST("/questionlikes", handlers.QuestionLikes)
	r.POST("/saveanswer", handlers.SaveAnswer)
	//answers
	r.POST("/acceptanswer", handlers.AcceptAnswer)
	r.POST("/answerlikes", handlers.AnswerLikes)
	r.POST("/answerdislikes", handlers.AnswerDisLikes)
	r.GET("/answeredit/:id", handlers.EditAnswer)
	r.POST("/answerupdate/:id", handlers.UpdateAnswer)
	r.POST("/answerdelete/:id", handlers.AnswerDelete)
	//tags
	r.GET("/tags/:id", handlers.Tags)
	r.GET("/tagedit/:id", handlers.EditTag)
	r.POST("/tagupdate", handlers.UpdateTag)
	r.GET("/categories", handlers.Categories)
	//users
	r.GET("/signup", handlers.SignUp)
	r.POST("/save", handlers.SaveUser)
	r.POST("/deleteuser/:id", handlers.DeleteUser)
	r.GET("/logout", handlers.Logout)
	r.GET("/login", handlers.Login)
	r.POST("/signin", handlers.SignIn)
	r.GET("/profile/:id", handlers.Profile)
	r.GET("/admin", handlers.Admin)
	r.GET("/rank", handlers.RankUser)
	//chat
	r.GET("/chat", handlers.Chat)

	m := melody.New()

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	return r
}