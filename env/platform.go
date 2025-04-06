package env

import (
	"encoding/json"
	"os"
	"runtime"
)

type Platform struct {
	OS       string `json:"os"`
	ARCH     string `json:"arch"`
	HOSTNAME string `json:"hostname"`
	CPUs     int    `json:"cpus"`
}

var instance *Platform

type MemoryProfile struct {
	Total uint64 `json:"total"`
	Heap  uint64 `json:"heapTotal"`
	Stack uint64 `json:"stackTotal"`
	GC    uint64 `json:"gcTotal"`
}

type HeapSnapshot struct {
	InUse    uint64 `json:"heapInUse"`
	IDLE     uint64 `json:"heapIdle"`
	Released uint64 `json:"heapReleased"`
	Objects  uint64 `json:"heapObjects"`
}

type StackSnapshot struct {
	InUse uint64 `json:"stackInUse"`
}

func GetPlatform() *Platform {
	if instance == nil {
		instance = &Platform{
			OS:   runtime.GOOS,
			ARCH: runtime.GOARCH,
			CPUs: runtime.NumCPU(),
		}
		instance.HOSTNAME, _ = os.Hostname()
	}
	return instance
}

func (platform *Platform) String() string {
	data, _ := json.Marshal(platform)
	return string(data)
}

func (platform *Platform) ActiveGoroutines() int {
	return runtime.NumGoroutine()
}

func (platform *Platform) MemoryProfile() *MemoryProfile {
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	return &MemoryProfile{
		Total: memStat.Sys,
		Heap:  memStat.HeapSys,
		Stack: memStat.StackSys,
		GC:    memStat.GCSys,
	}
}

func (profile *MemoryProfile) HeapSnapshot() *HeapSnapshot {
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	return &HeapSnapshot{
		InUse:    memStat.HeapInuse,
		IDLE:     memStat.HeapIdle,
		Released: memStat.HeapReleased,
		Objects:  memStat.HeapObjects,
	}
}

func (profile *MemoryProfile) StackSnapshot() *StackSnapshot {
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	return &StackSnapshot{
		InUse: memStat.StackInuse,
	}
}

func (profile *MemoryProfile) String() string {
	data, _ := json.Marshal(profile)
	return string(data)
}

func (snapshot *HeapSnapshot) String() string {
	data, _ := json.Marshal(snapshot)
	return string(data)
}

func (snapshot *StackSnapshot) String() string {
	data, _ := json.Marshal(snapshot)
	return string(data)
}
