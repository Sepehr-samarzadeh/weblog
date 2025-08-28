package main

import (
	"log"
	"net/http"
	"weblog/authentication"
	"weblog/db"
	"weblog/post"
	"weblog/userblog"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func AdddummyData(ctx *gin.Context) {
	var post post.BlogPost
	err := ctx.ShouldBindJSON(&post)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "cant bind with json"})
		return
	}

	err = post.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "cant save the post"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "post has been saved to db"})

}

func Register(ctx *gin.Context) {
	var user userblog.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "cant bind the values with json format"})
		return
	}
	err = user.RegisterUser()
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "cant register user"})
		return
	}

	token, err := authentication.GenerateJWTtoken(user.Email, int64(user.Id))
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "token is fucked!"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "user has been saved to db",
		"token": token,
	})

}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("no .env file found on system environment variables")
	}

	db.InitDB()

	router := gin.Default()

	router.LoadHTMLGlob("index.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{"message": "main page"})
	})

	router.POST("/data", AdddummyData)
	router.POST("register", Register)

	router.Run(":8080")

}
