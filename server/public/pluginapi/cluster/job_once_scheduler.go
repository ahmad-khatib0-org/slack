package cluster

import "sync"

// syncedCallback uses the mutex to make things predictable for the client: the
// callback will be called once at a time (the client does not need to worry
// about concurrency within the callback)
type syncedCallback struct {
	mu       sync.Mutex
	callback func(string, any)
}

type syncedJobs struct {
	mu   sync.RWMutex
	jobs map[string]*JobOnce
}

type JobOnceScheduler struct {
	pluginAPI JobPluginAPI

	startedMu sync.RWMutex
	started   bool

	activeJobs     *syncedJobs
	storedCallback *syncedCallback
}

var (
	schedulerOnce sync.Once
	s             *JobOnceScheduler
)
