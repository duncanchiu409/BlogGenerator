package databases

import (
	"blogAI/utils"
	"database/sql"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Blog struct {
	Id      string
	Title   string
	Content string
	Time    time.Time
}

type Blogs struct {
	mu *sync.Mutex
	db *sql.DB
}

const create string = `
  CREATE TABLE IF NOT EXISTS blogs (
  id BLOB NOT NULL PRIMARY KEY,
	title TEXT,
  content TEXT,
	time DATETIME NOT NULL
  );`

func (b *Blogs) Insert(blog *Blog) (string, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	utils.InfoLog.Println(blog.Time.String())
	_, err := b.db.Exec(`INSERT INTO blogs VALUES (?,?,?,?)`, blog.Id, blog.Title, blog.Content, blog.Time)
	if err != nil {
		return "", err
	}
	return blog.Id, nil
}

func (b *Blogs) Retreive(id string) (*Blog, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	row := b.db.QueryRow(`SELECT id, title, content, time FROM blogs WHERE id=?`, id)

	newBlog := Blog{}
	err := row.Scan(&newBlog.Id, &newBlog.Title, &newBlog.Content, &newBlog.Time)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &newBlog, nil
}

func (b *Blogs) Delete(id string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	_, err := b.db.Exec(`DELETE FROM blogs WHERE id=?`, id)
	if err != nil {
		return err
	}
	return nil
}

func NewBlogs(file string) (*Blogs, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	newBlogs := Blogs{}
	newBlogs.mu = &sync.Mutex{}
	newBlogs.db = db
	return &newBlogs, nil
}
