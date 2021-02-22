package internal

import (
	"argus/internal/config"
	"argus/internal/watcher"
	"errors"
)

type WatcherGateway struct {
	ConfigPath string
}

func (w WatcherGateway) Watch() error {
	configLoader := config.ConfigLoader{ConfigPath: w.ConfigPath}
	argusLoad, err := configLoader.Load()

	if(err!=nil){
		return errors.New("error loading loadedConfig")
	}

	for _, loadedConfig := range argusLoad.ConfigList {
		factory := watcher.WatcherFactory{}
		watcherCreated, err := factory.CreateWatcher(loadedConfig.Type)
		if(err==nil){
			watcherCreated.Watch(loadedConfig)
		}
	}
	return nil
}
