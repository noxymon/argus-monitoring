package internal

import (
	"testing"
	"time"
)

func TestWatcherGateway_Watch(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipped on short mode")
	}

	gateway := WatcherGateway{ConfigPath: "/home/noxymon/go-labs/argus/config.json"}
	err := gateway.Watch()
	if(err!=nil){
		t.Error("Failed : " + err.Error())
	}
	time.Sleep(5*time.Second)
}

