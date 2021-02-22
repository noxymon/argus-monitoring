package watcher

import (
	"argus/internal/config"
	"errors"
	"github.com/shirou/gopsutil/process"
	"os/exec"
	"strings"
)

type ProcessWatcher struct {

}

func (pw ProcessWatcher) Watch(config config.Config) error {
	processes, err := process.Processes()
	if(err!=nil){
		return errors.New("Error retrieving process")
	}

	alive := pw.isAlive(processes, config)
	if isNot(alive) {
		command := exec.Command(config.CommandLine)
		command.Start()
	}
	return nil
}

func (pw ProcessWatcher) isAlive(processes []*process.Process, config config.Config) bool {
	var alive = false
	for _, process := range processes {
		cmdline, _ := process.Cmdline()
		if cmdline != "" {
			match := strings.Contains(cmdline, config.CommandLine)
			if match {
				alive = true
			}
		}
	}
	return alive
}
