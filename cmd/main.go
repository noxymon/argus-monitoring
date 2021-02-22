package main

import (
	"flag"
	"github.com/go-co-op/gocron"
	"log"
	"os"
	"path/filepath"
	"time"
)



func main() {
	interval := flag.Int("interval", 2, "Interval of schedule check in seconds")
	//configPath := flag.String("config", currentPath()+"/config.json", "location of config path")
	//watcherGateway := internal.WatcherGateway{ConfigPath: *configPath}
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic("failed!!!")
	}
	scheduler := gocron.NewScheduler(location)
	scheduler.Every(interval).Second().Do(func() {
		log.Println(" Hello Schedule !")
	})
	scheduler.StartBlocking()
}

func startCron(interval *int) {
	log.Println("create cron schedule")
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic("failed!!!")
	}
	scheduler := gocron.NewScheduler(location)
	scheduler.Every(interval).Second().Do(showMessage)
	scheduler.StartBlocking()
}

func currentPath() string {
	currentPath, pathErr := filepath.Abs(filepath.Dir(os.Args[0]))
	if pathErr != nil {
		panic("failed get current path")
	}
	return currentPath
}

func showMessage() {
	log.Println(" Hello Schedule !")
}
