package main

import (
	"blogAI/databases"
	"blogAI/utils"
	"time"

	"github.com/google/uuid"
)

func main() {
	utils.InitCustomLogger("Database")

	const file string = "../databases/blogs.db"
	blogs, err := databases.NewBlogs(file)
	if err != nil {
		utils.ErrorLog.Fatal(err)
	}

	TestBlog := databases.Blog{
		Id:      uuid.New().String(),
		Title:   "A travel blog to Italy",
		Content: "Been to Museum & Beach",
		Time:    time.Now().UTC(),
	}
	blogs.Insert(&TestBlog)

	resultBlog, err := blogs.Retreive(TestBlog.Id)
	if err != nil {
		utils.ErrorLog.Fatal(err)
	}

	err = blogs.Delete("5f6730bb-5d75-4441-896f-c1705f35d73f")
	if err != nil {
		utils.ErrorLog.Println(err)
	}

	utils.InfoLog.Printf("id: %v", resultBlog.Id)
	utils.InfoLog.Printf("title: %v", resultBlog.Title)
	utils.InfoLog.Printf("content: %v", resultBlog.Content)
	utils.InfoLog.Printf("time: %v", resultBlog.Time)
}
