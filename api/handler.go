// +build linux

package api

import (
	"log"
	"strconv"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct{}

// GetDiskStats generates responce to a Ping request
func (s *Server) GetDiskStats(ctx context.Context, in *DiskStatus) (*DiskStatus, error) {
	log.Printf("Recieve message DISK %s", in.Diskstat)
	stat, err := disk.Usage("/")
	if err != nil {
		log.Fatal("stat read fail")
	}
	var status string

	status = " Total: " + strconv.FormatUint(stat.Total, 10) + " | Free: " + strconv.FormatUint(stat.Free, 10) + " | Used: " + strconv.FormatUint(stat.Used, 10)

	return &DiskStatus{Diskstat: status}, nil
}

// GetCpuStats generates responce to a Ping request
func (s *Server) GetCpuStats(ctx context.Context, in *CpuStatus) (*CpuStatus, error) {
	log.Printf("Recieve message CPU %s", in.Cpustat)
	stat, err := cpu.Info()
	if err != nil {
		log.Fatal("stat read fail")
	}
	var cpustatus string

	for _, s := range stat {
		cpustatus = "VendorID: " + s.VendorID + " | Model: " + s.Model + " | ModelName: " + s.ModelName
	}
	return &CpuStatus{Cpustat: cpustatus}, nil
}

// GetRamStats generates responce to a Ping request
func (s *Server) GetRamStats(ctx context.Context, in *RamStatus) (*RamStatus, error) {
	log.Printf("Recieve message RAM %s", in.Ramstat)
	stat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal("stat read fail")
	}
	var status string

	status = "Total: " + strconv.FormatUint(stat.Total, 10) + " | Used: " + strconv.FormatUint(stat.Used, 10) + " | Free: " + strconv.FormatUint(stat.Free, 10)

	return &RamStatus{Ramstat: status}, nil
}

// GetNetStats generates responce to a Ping request
func (s *Server) GetNetStats(ctx context.Context, in *NetStatus) (*NetStatus, error) {
	log.Printf("Recieve message NET %s", in.Netstat)
	stat, err := net.IOCounters(false)
	if err != nil {
		log.Fatal("stat read fail")
	}
	var status string

	for _, s := range stat {
		status = "Name: " + s.Name + " | BytesSent: " + strconv.FormatUint(s.BytesSent, 10) + " | BytesRecv: " + strconv.FormatUint(s.BytesRecv, 10)
	}
	return &NetStatus{Netstat: status}, nil
}
