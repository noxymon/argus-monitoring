package watcher

import "testing"

func TestWatcherFactory_CreateWatcherHttp(t *testing.T) {
	factory := WatcherFactory{}
	watcher, err := factory.CreateWatcher("http")
	if(err==nil){
		_, ok := watcher.(HttpWatcher)
		if isNot(ok) {
			t.Error("Type Assertion error")
		}
	}
}

func TestWatcherFactory_CreateWatcherProcess(t *testing.T) {
	factory := WatcherFactory{}
	watcher, err := factory.CreateWatcher("process")
	if(err==nil){
		_, ok := watcher.(ProcessWatcher)
		if isNot(ok) {
			t.Error("Type Assertion error")
		}
	}
}

func TestWatcherFactory_CreateWatcherUppercase(t *testing.T) {
	factory := WatcherFactory{}
	watcher, err := factory.CreateWatcher("PROCESS")
	if(err==nil){
		_, ok := watcher.(ProcessWatcher)
		if isNot(ok) {
			t.Error("Type Assertion error")
		}
	}
}