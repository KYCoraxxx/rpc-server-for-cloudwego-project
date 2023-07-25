package main

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"math/rand"
	"net"
	demo "rpc_server/kitex_gen/demo/studentservice"
	"strconv"
	"time"
)

var bindPort int

func main() {
	bindPort = rand.Intn(100) + 9000
	addr, _ := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(bindPort))

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	svr := demo.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr),
		server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "student",
		}), server.WithRegistryInfo(&registry.Info{
			Tags: map[string]string{
				"Cluster": "StudentCluster",
			}}), server.WithExitWaitTime(time.Minute))

	err = svr.Run()

	if err != nil {
		log.Fatal(err)
	}
}
