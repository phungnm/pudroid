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
		checkUrl,_ := models.GetShortenUrlByCode(sUrl.Code)
		if(*checkUrl==models.ShortenUrl{}){
						sUrl.Save()
			c.JSON(200,  gin.H{
				"message" : "Success",
				"ShortenUrl" :  sUrl,
			})
		} else {
			c.JSON(500, gin.H{"error": "Code already existed"})
		}

	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}
func GetShortenUrl(c*gin.Context){
	// type GetShortenUrl struct {
	// 	Code string `gorm:"unique" json:"Code"`
	// }
	// data := GetShortenUrl{}
		result,errs := models.GetShortenUrlByCode(c.Query("code"))
		fmt.Printf("%v\n", c.Query("code"))
		if(errs!=nil){
			if(*result!=models.ShortenUrl{}){
					c.JSON(200,gin.H{
			"shorten_url": result})
			}else {
					c.JSON(500, gin.H{"error": "Can't find this shorten_url"})
				}
			
		}else {
			c.JSON(500, gin.H{"error": "Something wrong here!"})
		}
}
