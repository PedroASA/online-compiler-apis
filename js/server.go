/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	
	"os/exec"
	"io/ioutil"

	"google.golang.org/grpc"
	pb "github.com/userName/otherModule/apis"
)

var (
	port = flag.Int("port", 8000, "The server port")
	in_file = "tmp/in.txt"
	code_file = "tmp/code.js"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedCodeServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Run(ctx context.Context, in *pb.CodeRequest) (*pb.CodeReply, error) {
	code := in.GetCode()
	input := in.GetInput()
	log.Printf("Received: %v\n %v", code, input)

	if err := ioutil.WriteFile(code_file, []byte(code), 0644); err != nil {
		return &pb.CodeReply{}, err
	}

	nodeCmd := exec.Command("node", code_file)



	// get scripts stdin, stdout, stderr
	cmdIn, _ := nodeCmd.StdinPipe()
	cmdOut, _ := nodeCmd.StdoutPipe()
	cmdErr, _ := nodeCmd.StderrPipe()

	// run script
	nodeCmd.Start()
	// write to its stdin
	if input != "" {
		cmdIn.Write([]byte(input))
	}
	cmdIn.Close()

	// Get its stdout, stderr
	stdout, _ := ioutil.ReadAll(cmdOut)
	stderr, _ := ioutil.ReadAll(cmdErr)

	// Wait command exit status.
	nodeCmd.Wait()

	// _, err = exec.Command("date", "-x").Output()
	// if err != nil {
	// 	switch e := err.(type) {
	// 	case *exec.Error:
	// 		fmt.Println("failed executing:", err)
	// 	case *exec.ExitError:
	// 		fmt.Println("command exit rc =", e.ExitCode())
	// 	default:
	// 		panic(err)
	// 	}
	// }

	return &pb.CodeReply{OutMessage: string(stdout), ErrMessage: string(stderr)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCodeServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
