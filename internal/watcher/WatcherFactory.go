package watcher

import (
	"errors"
	"strings"
)

type WatcherFactory struct {
}

func (w WatcherFactory) CreateWatcher(name string) (iWatcher, error) {
	switch strings.ToLower(name) {
	case "http":
		return HttpWatcher{}, nil
	case "process":
		return ProcessWatcher{}, nil
	}
	return nil, errors.New(name + " not recognize")
}
