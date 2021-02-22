package watcher

import (
	"argus/internal/config"
	"github.com/gojektech/heimdall/v6/httpclient"
	"os/exec"
)

type HttpWatcher struct{}

func (h HttpWatcher) Watch(config config.Config) error {
	alive := h.isAlive(config)

	if isNot(alive) {
		command := exec.Command(config.CommandLine)
		command.Start()
	}
	return nil
}

func (h HttpWatcher) isAlive(config config.Config) bool {
	var alive = false
	client := httpclient.NewClient()
	get, err := client.Get(config.HealthCheckURL, nil)
	if err != nil {
		alive = false
	}

	if get.StatusCode >= 200 && get.StatusCode <= 299 {
		alive = true
	}
	return alive
}