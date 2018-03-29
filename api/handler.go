package api

import (
	"log"
	"strconv"

	linuxproc "github.com/c9s/goprocinfo/linux"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct{}

// GetDiskStats generates responce to a Ping request
func (s *Server) GetDiskStats(ctx context.Context, in *DiskStatus) (*DiskStatus, error) {
	log.Printf("Recieve message DISK %s", in.Diskstat)
	stat, err := linuxproc.ReadDiskStats("/proc/diskstats")
	if err != nil {
		log.Fatal("stat read fail")
	}
	var status string

	for _, s := range stat {
		status = "Name: " + s.Name + " | ReadIOs: " + strconv.FormatUint(s.ReadIOs, 10) + " | ReadSectors: " + strconv.FormatUint(s.ReadSectors, 10) + " | WriteIOs: " + strconv.FormatUint(s.WriteIOs, 10) + " | WriteSectors: " + strconv.FormatUint(s.WriteSectors, 10)
	}
	return &DiskStatus{Diskstat: status}, nil
}

// GetCpuStats generates responce to a Ping request
func (s *Server) GetCpuStats(ctx context.Context, in *CpuStatus) (*CpuStatus, error) {
	log.Printf("Recieve message CPU %s", in.Cpustat)
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}
	var cpustatus string

	for _, s := range stat.CPUStats {
		cpustatus = "User: " + strconv.FormatUint(s.User, 10) + " | Nice: " + strconv.FormatUint(s.Nice, 10) + " | System: " + strconv.FormatUint(s.System, 10) + " | Idle: " + strconv.FormatUint(s.Idle, 10) + " | IOWait: " + strconv.FormatUint(s.IOWait, 10)
	}
	return &CpuStatus{Cpustat: cpustatus}, nil
}

// GetRamStats generates responce to a Ping request
func (s *Server) GetRamStats(ctx context.Context, in *RamStatus) (*RamStatus, error) {
	log.Printf("Recieve message RAM %s", in.Ramstat)
	stat, err := linuxproc.ReadVMStat("/proc/vmstat")
	if err != nil {
		log.Fatal("stat read fail")
	}
	var status string

	status = "PagePagein: " + strconv.FormatUint(stat.PagePagein, 10) + " | PagePageout: " + strconv.FormatUint(stat.PagePageout, 10) + " | PageSwapin: " + strconv.FormatUint(stat.PageSwapin, 10) + " | PageSwapout: " + strconv.FormatUint(stat.PageSwapout, 10) + " | PageFree: " + strconv.FormatUint(stat.PageFree, 10)

	return &RamStatus{Ramstat: status}, nil
}

// GetNetStats generates responce to a Ping request
func (s *Server) GetNetStats(ctx context.Context, in *NetStatus) (*NetStatus, error) {
	log.Printf("Recieve message NET %s", in.Netstat)
	stat, err := linuxproc.ReadNetStat("/proc/net/netstat")
	if err != nil {
		log.Fatal("stat read fail")
	}
	var status string

	status = "InMcastPkts: " + strconv.FormatUint(stat.InMcastPkts, 10) + " | OutMcastPkts: " + strconv.FormatUint(stat.OutMcastPkts, 10) + " | InOctets: " + strconv.FormatUint(stat.InOctets, 10) + " | OutOctets: " + strconv.FormatUint(stat.OutOctets, 10) + " | TCPTimeouts: " + strconv.FormatUint(stat.TCPTimeouts, 10)
	return &NetStatus{Netstat: status}, nil
}
