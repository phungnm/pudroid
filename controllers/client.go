package controllers

import(
	"github.com/gin-gonic/gin"
	"api/database"
	"fmt"
)
func Delete(c * gin.Context){
	db := database.DBConn()
    delete, err := db.Prepare("DELETE FROM posts WHERE id=?")
    if err != nil {
        panic(err.Error())
	}
	delete.Exec(c.Param("id"))
	c.JSON(200, gin.H{
		"messages": "deleted",
	})
    defer db.Close()
}
func Update(c * gin.Context){
	db := database.DBConn()
	type UpdatePost struct {
		Title string `form:"title" json:"title" binding:"required"`
		Body string `form:"body" json:"body" binding:"required"`
	}
	var json UpdatePost
	if err := c.ShouldBindJSON(&json); err == nil {
		edit, err := db.Prepare("UPDATE posts SET title=?, body=? WHERE id= " + c.Param("id"))
        if err != nil {
            panic(err.Error())
        }
		edit.Exec(json.Title, json.Body)
		
		c.JSON(200, gin.H{
			"messages": "edited",
		})
	}else{
		c.JSON(500, gin.H{"error": err.Error()})
	}
    defer db.Close()
}
type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"body"`
}
func Create(c * gin.Context){
	db := database.DBConn()
	type CreatePost struct {
		Title string `form:"title" json:"title" binding:"required"`
		Body string `form:"body" json:"body" binding:"required"`
	}

	var json CreatePost

	if err := c.ShouldBindJSON(&json); err == nil {
		insPost, err := db.Prepare("INSERT INTO posts(title, body) VALUES(?,?)",)
		if err != nil {
			c.JSON(500, gin.H{
				"messages" : err,
			})
		}

		insPost.Exec(json.Title, json.Body) 
		c.JSON(200, gin.H{
			"messages": "inserted",
		})

	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer db.Close()
}
func Read(c * gin.Context){

	db := database.DBConn()
	rows, err := db.Query("SELECT * FROM posts WHERE id = " + c.Param("id"))
	if err != nil{
		 fmt.Println(err)
		c.JSON(500, gin.H{
			"messages" : "Story not found",
		});
	} else {
			post := Post{}
	for rows.Next(){
		var id int
		var title, body string
		err = rows.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		post.Id = id
		post.Title = title
		post.Content = body
	}

	c.JSON(200, post)
	}

	defer db.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
func GetAll(c * gin.Context){

	db := database.DBConn()
	rows, err := db.Query("SELECT * FROM posts WHERE 1")
	if err != nil{
		 fmt.Println(err)
		c.JSON(500, gin.H{
			"messages" : "Story not found",
		});
	} else {
	posts := []Post{}
	for rows.Next(){
		post := Post{}
		var id int
		var title, body string
		err = rows.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		post.Id = id
		post.Title = title
		post.Content = body
		posts = append(posts,post)

	}

	c.JSON(200, posts)
	}

	defer db.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}