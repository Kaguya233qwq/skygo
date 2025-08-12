package plugins

import (
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"

	logger "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	status()
	logger.Info("Plugin:Status loaded successfully")
}

func formatBytes(bytes uint64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
	)

	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/TB)
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d bytes", bytes)
	}
}

type SystemInfo struct {
	OS                string
	Architecture      string
	Platform          string
	CPUUsage          float64
	MemoryTotal       string
	MemoryAvailable   string
	MemoryUsed        string
	MemoryUsedPercent string
	Uptime            string
	ProcessPID        int32
	ProcessRSS        string
	ProcessVMS        string
}

func getSystemInfo() (SystemInfo, error) {
	hi, err := host.Info()
	if err != nil {
		logger.Errorf("Error getting host information: %v\n", err)
		return SystemInfo{}, err
	}
	fmt.Printf("Hostname: %s\n", hi.Hostname)
	fmt.Printf("OS: %s\n", hi.OS)
	fmt.Printf("Platform: %s\n", hi.Platform)
	fmt.Printf("Platform Family: %s\n", hi.PlatformFamily)
	fmt.Printf("Platform Version: %s\n", hi.PlatformVersion) 
	fmt.Printf("Kernel Version: %s\n", hi.KernelVersion)
	fmt.Printf("Architecture: %s\n", hi.KernelArch)

	// 获取当前设备 CPU 占用
	fmt.Println("\n--- CPU Usage ---")
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("Error getting CPU percent: %v\n", err)
	} else {
		if len(percent) > 0 {
			fmt.Printf("Overall CPU Usage: %.2f%%\n", percent[0])
		}
	}

	// 获取当前设备内存占用
	fmt.Println("\n--- Memory Usage (Overall) ---")
	vmem, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error getting virtual memory info: %v\n", err)
	} else {
		fmt.Printf("Total Memory: %s\n", formatBytes(vmem.Total))
		fmt.Printf("Available Memory: %s\n", formatBytes(vmem.Available))
		fmt.Printf("Used Memory: %s\n", formatBytes(vmem.Used))
		fmt.Printf("Used Percent: %.2f%%\n", vmem.UsedPercent)
	}

	// 获取开机时间
	fmt.Println("\n--- System Uptime ---")
	var uptimeDuration time.Duration
	uptimeSeconds, err := host.Uptime()
	if err != nil {
		fmt.Printf("Error getting uptime: %v\n", err)
	} else {
		uptimeDuration = time.Duration(uptimeSeconds) * time.Second
		fmt.Printf("System Uptime: %s\n", uptimeDuration.String())
	}

	// 获取当前程序内存占用量
	fmt.Println("\n--- Current Process Memory Usage ---")
	pid := os.Getpid() // Get the current process ID
	p, err := process.NewProcess(int32(pid))
	var memInfo *process.MemoryInfoStat
	if err != nil {
		fmt.Printf("Error creating process object for PID %d: %v\n", pid, err)
	} else {
		memInfo, err = p.MemoryInfo()
		if err != nil {
			fmt.Printf("Error getting memory info for process %d: %v\n", pid, err)
		} else {
			fmt.Printf("Process PID: %d\n", pid)
			fmt.Printf("Process RSS (Resident Set Size): %s\n", formatBytes(memInfo.RSS))
			fmt.Printf("Process VMS (Virtual Memory Size): %s\n", formatBytes(memInfo.VMS))
		}
	}

	fmt.Println("\n--- End of Information ---")
	return SystemInfo{
		OS:                hi.OS,
		Architecture:      hi.KernelArch,
		Platform:          hi.Platform,
		CPUUsage:          percent[0],
		MemoryTotal:       formatBytes(vmem.Total),
		MemoryAvailable:   formatBytes(vmem.Available),
		MemoryUsed:        formatBytes(vmem.Used),
		MemoryUsedPercent: fmt.Sprintf("%.2f%%\n", vmem.UsedPercent),
		Uptime:            uptimeDuration.String(),
		ProcessPID:        int32(pid),
		ProcessRSS:        formatBytes(memInfo.RSS),
		ProcessVMS:        formatBytes(memInfo.VMS),
	}, nil
}

func status() {
	zero.OnCommand("status").Handle(func(ctx *zero.Ctx) {
		systemInfo, err := getSystemInfo()
		if err != nil {
			logger.Errorf("Failed to get system info: %v", err)
			ctx.Send(message.Text("Failed to retrieve system information."))
		}
		ctx.Send(message.Text("System status：\n" +
			"Platform: " + systemInfo.Platform + "\n" +
			"Architecture: " + systemInfo.Architecture + "\n" +
			"OS: " + systemInfo.OS + "\n" +
			"CPU: " + fmt.Sprintf("%.2f%%\n", systemInfo.CPUUsage) +
			"MEM: " + systemInfo.MemoryUsed + "/" + systemInfo.MemoryTotal +
			fmt.Sprintf("(%s)", systemInfo.MemoryUsedPercent) + "\n" +
			"BootTime: " + systemInfo.Uptime + "\n" +
			"PID: " + fmt.Sprintf("%d\n", systemInfo.ProcessPID) +
			"ProcessRSS: " + systemInfo.ProcessRSS + "\n" +
			"ProcessVMS: " + systemInfo.ProcessVMS,
		))
	})
}
