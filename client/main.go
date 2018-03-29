package main

import (
	"context"
	"log"
	"monitoring/api"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewInfoStatusClient(conn)

	diskresponse, err := c.GetDiskStats(context.Background(), &api.DiskStatus{Diskstat: "disk stat ok"})
	if err != nil {
		log.Fatalf("Error when calling GetDiskStats: %s", err)
	}
	log.Printf("Response from server: %s", diskresponse.Diskstat)

	cpuresponse, err := c.GetCpuStats(context.Background(), &api.CpuStatus{Cpustat: "cpu stat ok"})
	if err != nil {
		log.Fatalf("Error when calling GetCpuStats: %s", err)
	}
	log.Printf("Response from server: %s", cpuresponse.Cpustat)

	ramresponse, err := c.GetRamStats(context.Background(), &api.RamStatus{Ramstat: "ram stat ok"})
	if err != nil {
		log.Fatalf("Error when calling GetRamStats: %s", err)
	}
	log.Printf("Response from server: %s", ramresponse.Ramstat)

	netresponse, err := c.GetNetStats(context.Background(), &api.NetStatus{Netstat: "net stat ok"})
	if err != nil {
		log.Fatalf("Error when calling GetNetStats: %s", err)
	}
	log.Printf("Response from server: %s", netresponse.Netstat)
}
