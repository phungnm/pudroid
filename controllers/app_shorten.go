package controllers

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"pudroid/models"
	"net/http"
)

func GetShorten(c*gin.Context){
		code := c.Param("code")
		result,errs := models.GetShortenUrl(map[string]interface{}{ "code": code})

		if (errs == nil) {
			if(*result!=models.ShortenUrl{}){
				c.Redirect(http.StatusMovedPermanently, result.Url)
			}else {
					c.JSON(http.StatusOK, gin.H{"status":0,"error": "Can't find this shorten_url"})
				}
			
		}else {
				fmt.Println(errs)
			  c.JSON(http.StatusOK, gin.H{"status":0,"error": "Error"})
		}
}
