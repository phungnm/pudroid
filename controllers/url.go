package controllers

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"pudroid/models"
)
func HelloUrl(c * gin.Context){
	c.JSON(200,  gin.H{
			"message" : "Hello This is Pudroid Url"})
}
func AddShortenUrl(c * gin.Context){
	sUrl := models.ShortenUrl{}

	if err := c.ShouldBindJSON(&sUrl); err == nil {
		if sUrl.Code!=""{
			checkUrl,_ := models.GetShortenUrlByCode(sUrl.Code)
			if (*checkUrl==models.ShortenUrl{}) {

							sUrl.Create()
				c.JSON(200,  gin.H{
					"message" : "Success",
					"ShortenUrl" :  sUrl,
				})
			} else {
				c.JSON(500, gin.H{"error": "Code already existed"})
			}
		} else{
			sUrl.Create()
			sUrl.Code = fmt.Sprint(sUrl.ID)
			sUrl.Update()
			c.JSON(200,  gin.H{
					"message" : "Success",
					"ShortenUrl" :  sUrl,
				})
		}


	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}
func GetShortenUrl(c*gin.Context){
		code := c.Query("code")
		result,errs := models.GetShortenUrlByCode(code)

		if (errs == nil) {
			if(*result!=models.ShortenUrl{}){
					c.JSON(200,gin.H{
			"shorten_url": result})
			}else {
					c.JSON(500, gin.H{"error": "Can't find this shorten_url"})
				}
			
		}else {
			  c.JSON(500, gin.H{"error": "Error"})
		}
}
