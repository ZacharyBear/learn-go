package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	setupPing(router)
	setupAsciiJSON(router)
	renderHTML(router)
	go serverPush()
	JSONP(router)
	formBinding(router)

	// router.Run() // listen and serve on 0.0.0.0:8080
	router.Run("localhost:8080")
}

func setupPing(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func setupAsciiJSON(router *gin.Engine) {
	router.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO 语言",
			"tag":  "<br>",
		}

		// Output {"lang":"GO \u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})
}

func renderHTML(router *gin.Engine) {
	// router.LoadHTMLFiles("templates/me.html")
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "me.html", gin.H{
			"title": "Bonjour! Je suis Zenkie Bear.",
		})
	})
	router.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "Users",
		})
	})
}

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
	<script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color: red;">Bienvenue, Ginner!</h1>
</html>
`))

func serverPush() {
	router := gin.New()
	router.Static("/assets", "./assets")
	router.SetHTMLTemplate(html)

	router.GET("/", func(ctx *gin.Context) {
		if pusher := ctx.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		ctx.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})
	router.RunTLS("localhost:8081", "ssl.pem", "ssl.key")
}

func JSONP(router *gin.Engine) {
	router.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// Request: /jsonp?callback=x
		// Output: x({"foo": "bar"})
		c.JSONP(http.StatusOK, data)
	})
	router.RunTLS("localhost:8081", "ssl.pem", "ssl.key")
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func formBinding(router *gin.Engine) {
	// curl -v --form user=john --form password=doe http://localhost:8080/login
	router.POST("/login", func(ctx *gin.Context) {
		var form LoginForm
		// Explicit binding
		// ctx.ShouldBindWith(&form, binding.Form)
		if ctx.ShouldBind(&form) == nil {
			if form.User == "john" && form.Password == "doe" {
				ctx.JSON(http.StatusOK, gin.H{"status": "You are logged in!"})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
		log.Print(form)
	})
}
