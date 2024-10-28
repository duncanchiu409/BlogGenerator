package workers

import (
	"blogAI/types"
	"blogAI/utils"
	"log"
	"net/rpc"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CallGetTask() (*types.GetTaskReply, error) {
	args := types.GetTaskArgs{}
	args.Msg = "Hello World 🎃"
	reply := types.GetTaskReply{}
	s := call("Master.GetTask", &args, &reply)
	if s.Code() != codes.OK {
		return nil, s.Err()
	}
	return &reply, nil
}

func call(rpcname string, args interface{}, reply interface{}) *status.Status {
	sockname := utils.CoordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	s, ok := status.FromError(err)
	if !ok {
		log.Fatalf("rpcname: %v", s.Err())
	}
	return s
}

func NewWorker() {
	for {
		reply, err := CallGetTask()
		if err != nil {
			log.Printf("CallGetTask failed: %v", err)
		}
		switch reply.TaskType {
		case types.GenerateContentTaskType:
			log.Printf("tasktype %v", reply.TaskType)
		case types.GeneratePictureTaskType:
			log.Printf("tasktype %v", reply.TaskType)
		}
		time.Sleep(time.Second * time.Duration(10))
	}
}
