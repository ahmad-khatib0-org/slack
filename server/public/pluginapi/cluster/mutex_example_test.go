package cluster_test

import (
	"github.com/ahmad-khatib0-org/slack/server/public/plugin"
	"github.com/ahmad-khatib0-org/slack/server/public/pluginapi/cluster"
)

//nolint:staticcheck
func ExampleMutex() {
	// Use p.API from your plugin instead.
	pluginAPI := plugin.API(nil)

	m, err := cluster.NewMutex(pluginAPI, "key")
	if err != nil {
		panic(err)
	}
	m.Lock()
	// critical section
	m.Unlock()
}
