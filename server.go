package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Unix()

		// RAM metrikalari
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		// CPU foizi
		percent, _ := cpu.Percent(0, false) // 0 = instant foiz, false = umumiy CPU
		cpuUsage := 0.0
		if len(percent) > 0 {
			cpuUsage = percent[0]
		}

		// PID
		pid := os.Getpid()

		// Prometheus formatidagi metrikalar
		fmt.Fprintf(w, "app_up 1\n")
		fmt.Fprintf(w, "app_time %d\n", now)
		fmt.Fprintf(w, "app_pid %d\n", pid)
		fmt.Fprintf(w, "app_alloc_bytes %d\n", m.Alloc)
		fmt.Fprintf(w, "app_total_alloc_bytes %d\n", m.TotalAlloc)
		fmt.Fprintf(w, "app_sys_bytes %d\n", m.Sys)
		fmt.Fprintf(w, "app_num_gc %d\n", m.NumGC)
		fmt.Fprintf(w, "app_cpu_percent %.2f\n", cpuUsage) // CPU foizda
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from ", r.Host)
	})

	port := ":8080"
	fmt.Println("Server running on", port)
	http.ListenAndServe(port, nil)
}
