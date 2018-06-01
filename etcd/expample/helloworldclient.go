package main

import (
	"flag"
	"fmt"
	"time"

	"strconv"

	grpclb "grpcdemo/etcd/balance"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpcdemo/etcd/expample/pb"
)

var (
	serv = flag.String("service", "hello_service", "service name")
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()
	fmt.Println("serv", *serv)
	r := grpclb.NewResolver(*serv)
	b := grpc.RoundRobin(r)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b))

	if err != nil {
		panic(err)
	}
	fmt.Println("conn...")
	client := pb.NewGreeterClient(conn)
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
		if err == nil {
			fmt.Printf("%v: Reply is %s\n", t, resp.Message)
		} else {
			fmt.Println(err)
		}
	}
}