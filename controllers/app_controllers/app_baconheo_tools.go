package app_controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"pudroid/helpers"
	"github.com/lithammer/shortuuid"
	"os"
	"path/filepath"
	"pudroid/config"
	"strconv"
)

func ConvertColumnComma(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"baconheo/convertColumnComma", gin.H{"menu_baconheo":true,"convertColumnComma":true,"title": "Convert column comma","extra_js": []string{"convertCommaColumn.js"} })
}
func SubmitSplitExcel(ctx *gin.Context){
	dir, _ := os.Getwd()
	ticket := shortuuid.New()
	savePath:=dir+"/tickets/"+ticket
	os.MkdirAll(savePath,0777)

	rows,_:= strconv.Atoi(ctx.PostForm("rows"))

	keep_header,_:= strconv.Atoi(ctx.PostForm("keep_header"))
	file, err := ctx.FormFile("file")
	if err!=nil {
		panic(err.Error())
	}

	filename := filepath.Base(file.Filename)
	filePath := savePath+"/"+filename
	if err := ctx.SaveUploadedFile(file,filePath); err != nil {
		panic(err.Error())
	}

	helpers.SplitExcel(filePath,rows,dir+"/excel_data/"+ticket,keep_header);

	helpers.ZipFolder(dir+"/excel_data/"+ticket,savePath,filename+".zip");

	url := config.Config.Base_URL+"/tickets/"+ticket+"/"+filename+".zip"
	ctx.JSON(http.StatusOK, gin.H{ "status": 1,"url":url})
}
func SplitExcel(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"baconheo/splitExcel", gin.H{"menu_baconheo":true,"splitExcel":true,"title": "Split Excel files","extra_js": []string{"splitExcel.js"} })
}
func CheckIP(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"baconheo/checkIP", gin.H{"menu_baconheo":true,"checkIP":true,"title": "Check IP",
	"extra_css": []string{"datatable/datatables.min.css","datatable/buttons.dataTables.min.css","datatable/dataTables.searchHighlight.css"},
	"extra_js": []string{"datatable/jquery.highlight.js","datatable/datatables.min.js","datatable/dataTables.buttons.min.js","datatable/dataTables.searchHighlight.min.js","datatable/buttons.flash.min.js","datatable/buttons.html5.min.js","datatable/buttons.print.min.js","xlsx.full.min.js","checkIP.js"} })
}
func QueryTable(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"baconheo/queryTable", gin.H{"menu_baconheo":true,"queryTable":true,"title": "Query Table",
	"extra_css": []string{"datatable/datatables.min.css","datatable/buttons.dataTables.min.css","datatable/dataTables.searchHighlight.css"},
	"extra_js": []string{"datatable/jquery.highlight.js","datatable/datatables.min.js","datatable/dataTables.buttons.min.js","datatable/dataTables.searchHighlight.min.js","datatable/buttons.flash.min.js","datatable/buttons.html5.min.js","datatable/buttons.print.min.js","xlsx.full.min.js","queryTable.js"} })
}
