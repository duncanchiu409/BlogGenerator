package main

import (
	"blogAI/master"
	"blogAI/server"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Request struct {
	method    string
	url       string
	startTime time.Time
}

type LinearizabilityChecker struct {
	requestMap map[string]*Request
	cond       *sync.Cond
}

func NewLinearizabilityChecker() *LinearizabilityChecker {
	l := LinearizabilityChecker{}
	l.requestMap = map[string]*Request{}
	mu := sync.Mutex{}
	l.cond = sync.NewCond(&mu)
	go l.rescheduler()
	return &l
}

func (l *LinearizabilityChecker) CheckRequest(c *gin.Context) {
	id := c.GetString("taskid")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	val, ok := l.requestMap[id]
	if ok {
		if val.method == c.Request.Method && val.url == c.Request.URL.String() {
			c.AbortWithStatus(http.StatusTooManyRequests)
		}
	}
	newRequest := Request{}
	newRequest.method = c.Request.Method
	newRequest.url = c.Request.URL.String()
	newRequest.startTime = time.Now().UTC()
	l.requestMap[id] = &newRequest
	c.Next()
}

func (l *LinearizabilityChecker) rescheduler() {
	for {
		l.cond.L.Lock()
		defer l.cond.L.Unlock()
		for id, request := range l.requestMap {
			currTime := time.Now().UTC()
			if currTime.Sub(request.startTime).Minutes() >= 1 {
				delete(l.requestMap, id)
				l.cond.Broadcast()
			}
		}
	}
}

func main() {
	go func() {
		master.InitCustomLogger()
		m := master.NewMaster()
		for m.Done() == false {
			time.Sleep(time.Second)
		}
	}()

	// c := NewLinearizabilityChecker()
	router := gin.Default()
	router.GET("/test", server.Test)
	// router.Use(c.CheckRequest)
	router.POST("/create", server.CreateTask)
	router.Run()
}
