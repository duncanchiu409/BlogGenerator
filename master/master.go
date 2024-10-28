package master

import (
	"blogAI/types"
	"blogAI/utils"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
	"time"
)

type Task struct {
	id        string
	startTime time.Time
	endTime   time.Time
	status    State
}

type Master struct {
	tasks      map[string]*Task
	tRemaining int
	cond       *sync.Cond
}

func (m *Master) CreateTask(args *types.CreateTaskArgs, reply *types.CreateTaskReply) error {
	m.cond.L.Lock()
	defer m.cond.L.Unlock()
	newTask := Task{}
	newTask.id = args.TaskId
	newTask.startTime = time.Now().UTC()
	newTask.status = unstarted
	m.tasks[args.TaskId] = &newTask
	m.tRemaining += 1
	return nil
}

func (m *Master) GetTask(args *types.GetTaskArgs, reply *types.GetTaskReply) error {
	m.cond.L.Lock()
	defer m.cond.L.Unlock()
	InfoLog.Println(args.id)
	reply.TaskType = types.GenerateContentTaskType
	return nil
}

func (m *Master) server() {
	rpc.Register(m)
	rpc.HandleHTTP()
	sockname := utils.CoordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		ErrorLog.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func (m *Master) Done() bool {
	return false
}

func NewMaster() *Master {
	m := Master{}
	m.tasks = map[string]*Task{}
	mu := sync.Mutex{}
	m.cond = sync.NewCond(&mu)
	m.tRemaining = 0
	m.server()
	return &m
}
