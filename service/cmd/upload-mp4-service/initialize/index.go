package initialize

import (
	jobapp "app/cmd/upload-mp4-service/job"
	"app/cmd/upload-mp4-service/queue"
	"sync"
)

func Run() {
	list := []func(){
		runHttpSrver,
		jobapp.InitJob,
		queue.InitQueue,
	}

	var wg sync.WaitGroup

	for _, f := range list {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(f)
	}

	wg.Wait()
}
