package watcher

import "argus/internal/config"

type iWatcher interface {
	Watch(config config.Config) error
}

func isNot(alive bool) bool {
	return !alive
}