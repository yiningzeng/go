package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println("Hello Go!")
	os := runtime.GOOS

	if os == "windows" {
		fmt.Println("hello win")
	} else if os == "linux" {
		fmt.Println("hello linux")
	}

	connStr := "postgres://postgres:baymin1024@192.168.31.75/power_ai?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, project_id FROM train_record WHERE project_id = 'hMohgb9RA'")
	var id int
	var project_id string
	for rows.Next() { //作为循环条件来迭代获取结果集Rows
		err = rows.Scan(&id, &project_id) //1 user1update computing 2019-02-20
		fmt.Println(id, project_id)
	}
	router := gin.Default()
	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
