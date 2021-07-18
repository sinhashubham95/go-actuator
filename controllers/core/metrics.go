package core

import (
	"github.com/sinhashubham95/go-actuator/models"
	"runtime"
)

// GetMetrics is the used to get the runtime metrics
func GetMetrics() *models.MemStats {
	return getRuntimeMetrics()
}

func getRuntimeMetrics() *models.MemStats {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	bySize := make([]models.BySizeElement, 0, len(memStats.BySize))
	for _, size := range memStats.BySize {
		bySize = append(bySize, models.BySizeElement{
			Size:         size.Size,
			MAllocations: size.Mallocs,
			Frees:        size.Frees,
		})
	}

	return &models.MemStats{
		Alloc:         memStats.Alloc,
		TotalAlloc:    memStats.TotalAlloc,
		Sys:           memStats.Sys,
		Lookups:       memStats.Lookups,
		MAllocations:  memStats.Mallocs,
		Frees:         memStats.Frees,
		HeapAlloc:     memStats.HeapAlloc,
		HeapSys:       memStats.HeapSys,
		HeapIdle:      memStats.HeapIdle,
		HeapInuse:     memStats.HeapInuse,
		HeapReleased:  memStats.HeapReleased,
		HeapObjects:   memStats.HeapObjects,
		StackInuse:    memStats.StackInuse,
		StackSys:      memStats.StackSys,
		MSpanInuse:    memStats.MSpanInuse,
		MSpanSys:      memStats.MSpanSys,
		MCacheInuse:   memStats.MCacheInuse,
		MCacheSys:     memStats.MCacheSys,
		BuckHashSys:   memStats.BuckHashSys,
		GCSys:         memStats.GCSys,
		OtherSys:      memStats.OtherSys,
		NextGC:        memStats.NextGC,
		LastGC:        memStats.LastGC,
		PauseTotalNs:  memStats.PauseTotalNs,
		PauseNs:       memStats.PauseNs,
		PauseEnd:      memStats.PauseEnd,
		NumGC:         memStats.NumGC,
		NumForcedGC:   memStats.NumForcedGC,
		GCCPUFraction: memStats.GCCPUFraction,
		EnableGC:      memStats.EnableGC,
		DebugGC:       memStats.DebugGC,
		BySize:        bySize,
	}
}
