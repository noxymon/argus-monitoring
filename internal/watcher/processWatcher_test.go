package watcher

import (
	"argus/internal/config"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"testing"
	"time"
)

func TestProcessList(t *testing.T) {
	skipOnShort(t)
	processes, err := process.Processes()

	if(err!=nil){
		t.Error("error !")
	}

	for _,s:= range processes{
		cmdline, _ := s.Cmdline()
		if(cmdline != ""){
			fmt.Println(cmdline)
		}
	}
}

func TestProcessWatcher_Watch(t *testing.T) {
	skipOnShort(t)
	c := config.Config{
		Type:           "process",
		HealthCheckURL: "",
		ProcessName:    "gedit",
		CommandLine:    "gedit",
	}
	watcher := ProcessWatcher{}
	watcher.Watch(c)
	time.Sleep(1 * time.Second)
}

func skipOnShort(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test in short mode")
	}
}
