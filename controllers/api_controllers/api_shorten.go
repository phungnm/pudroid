package api_controllers

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"pudroid/models"
	"net/http"
	"github.com/jinzhu/gorm"
)
func HelloUrl(c * gin.Context){
	c.JSON(200,  gin.H{ "status": 1,
			"message" : "Hello This is Pudroid Url"})
}
func AddShortenAPI(c * gin.Context){
	sUrl := models.ShortenUrl{}

	if err := c.ShouldBindJSON(&sUrl); err == nil {

		if sUrl.Code!=""{
			checkUrl,_ := models.GetShortenUrl(map[string]interface{}{ "code": sUrl.Code})
			if (*checkUrl==models.ShortenUrl{}) {
							sUrl.Create()
				c.JSON(http.StatusOK,  gin.H{ "status": 1,

					"message" : "Success",
					"ShortenUrl" :  sUrl,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{ "status": 0,"error": "Custom code already existed"})
			}
		} else{
			sUrl.Create()
			sUrl.Code = fmt.Sprint(sUrl.ID)
			sUrl.Update()

			c.JSON(200,  gin.H{ "status": 1,
					"message" : "Success",
					"ShortenUrl" :  sUrl,
				})
		}


	} else {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{ "status": 1,"error": err.Error()})
	}
}
func GetShortenAPI(c*gin.Context){
		code := c.Query("code")
		result,errs := models.GetShortenUrl(map[string]interface{}{ "code": code})

		if (errs == nil) {
			c.JSON(http.StatusBadRequest,gin.H{ "status": 1,
			"shorten_url": result})
					
			
		}else {
			if(gorm.IsRecordNotFoundError(errs)){
					c.JSON(http.StatusNotFound, gin.H{ "status": 1,"error": "Can't find this shorten_url"})
			}else {
				fmt.Println(errs)
			  c.JSON(http.StatusBadRequest, gin.H{ "status": 1,"error": "Error"})
			}
		}
}
