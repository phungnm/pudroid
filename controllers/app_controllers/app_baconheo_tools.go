package app_controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"pudroid/helpers"
)

func ConvertColumnComma(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"baconheo/convertColumnComma", gin.H{"menu_baconheo":true,"convertColumnComma":true,"title": "Convert column comma","extra_js": []string{"convertCommaColumn.js"} })
}
func SubmitSplitExcel(ctx *gin.Context){
	helpers.Test();
	time.Sleep(10 * time.Second)
	ctx.HTML(http.StatusOK,"baconheo/convertColumnComma", gin.H{"menu_baconheo":true,"convertColumnComma":true,"title": "Convert column comma","extra_js": []string{"convertCommaColumn.js"} })

}
func SplitExcel(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"baconheo/splitExcel", gin.H{"menu_baconheo":true,"splitExcel":true,"title": "Split Excel files","extra_js": []string{"splitExcel.js"} })
}
func CheckIP(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"baconheo/checkIP", gin.H{"menu_baconheo":true,"checkIP":true,"title": "Check IP",
	"extra_css": []string{"datatable/datatables.min.css","datatable/buttons.dataTables.min.css","datatable/dataTables.searchHighlight.css"},
	"extra_js": []string{"datatable/jquery.highlight.js","datatable/datatables.min.js","datatable/dataTables.buttons.min.js","datatable/dataTables.searchHighlight.min.js","datatable/buttons.flash.min.js","datatable/buttons.html5.min.js","datatable/buttons.print.min.js","xlsx.full.min.js","checkIP.js"} })
}

