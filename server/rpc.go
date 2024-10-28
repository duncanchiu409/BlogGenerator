package server

import (
	"blogAI/types"
	"blogAI/utils"
	"log"
	"net/rpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CallCreateTask(content string) error {
	args := types.CreateTaskArgs{}
	args.Content = content
	reply := types.GetTaskReply{}

	s := call("Master.CreateTask", &args, &reply)
	if s.Code() != codes.OK {
		return s.Err()
	}
	return nil
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
