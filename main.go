package main

import (
	// "fmt"
	//"pudroid/database"
	"github.com/gin-gonic/gin"
	 "pudroid/controllers/app_controllers"
	"pudroid/config"
	"github.com/foolin/gin-template"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/thoas/go-funk"
	"strconv"
	"github.com/itsjamie/gin-cors"
	"time"
)
func AuthenticationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User needs to be signed in to access this service"})
			c.Abort()
			return
		}
		if len(auths) != 0 {
			authType := session.Get("authType")
			if authType == nil || !funk.ContainsString(auths, authType.(string)) {
				c.JSON(http.StatusForbidden, gin.H{"error": "invalid request, restricted endpoint"})
				c.Abort()
				return
			}
		}
		// add session verification here, like checking if the user and authType
		// combination actually exists if necessary. Try adding caching this (redis)
		// since this middleware might be called a lot
		c.Next()
	}
}
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/public", "./public")
	router.Static("/assets", "./assets")
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views/",
		Extension: ".html",
		Master:    "layouts/master",
		DisableCache: true,
	})

	//ROUTER
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK,"index", gin.H{"title": "Home","menu_home": true,"extra_js": []string{} })
	})	
	tools := router.Group("/tools")
	{
		tools.GET("/convertColumnComma", app_controllers.ConvertColumnComma)
		tools.GET("/splitExcel", app_controllers.SplitExcel)
		tools.GET("/checkIP", app_controllers.CheckIP)
		tools.POST("/splitExcel", app_controllers.SubmitSplitExcel)


		
	}


	return router
}

func main() {
  	router := setupRouter()
  	router.Use(cors.Middleware(cors.Config{
	Origins:        "*",
	Methods:        "GET, PUT, POST, DELETE",
	RequestHeaders: "Origin, Authorization, Content-Type",
	ExposedHeaders: "",
	MaxAge: 50 * time.Second,
	Credentials: true,
	ValidateHeaders: false,
}))
	router.Run(":"+strconv.Itoa(config.Config.Port)) // Ứng dụng chạy tại cổng 3000
}