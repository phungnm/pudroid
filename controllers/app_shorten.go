package controllers

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"pudroid/models"
	"net/http"
)

func GetShorten(c*gin.Context){
		code := c.Param("code")
		result,errs := models.GetShortenAPIByCode(code)

		if (errs == nil) {
			if(*result!=models.ShortenUrl{}){
				c.Redirect(http.StatusMovedPermanently, result.Url)
			}else {
					c.JSON(500, gin.H{"error": "Can't find this shorten_url"})
				}
			
		}else {
				fmt.Println(errs)
			  c.JSON(500, gin.H{"error": "Error"})
		}
}
