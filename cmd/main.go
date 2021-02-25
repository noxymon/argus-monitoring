package main

import (
	"argus/internal"
	"flag"
	"github.com/go-co-op/gocron"
	"os"
	"path/filepath"
	"time"
)

func main() {
	interval := *flag.Int("interval", 2, "Interval of schedule check in seconds")

	configPath := *flag.String("config", currentPath(), "location of config path")
	watcherGateway := internal.WatcherGateway{ConfigPath: configPath}

	scheduler := createScheduler()
	scheduler.Every(interval).Second().Do(watcherGateway.Watch)
	scheduler.StartBlocking()
}

func createScheduler() *gocron.Scheduler {
	location, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := gocron.NewScheduler(location)
	return scheduler
}

func currentPath() string {
	currentPath, pathErr := filepath.Abs(filepath.Dir(os.Args[0]))
	if pathErr != nil {
		panic("failed get current path")
	}
	return currentPath
}
