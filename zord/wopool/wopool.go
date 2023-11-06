package wopool

import (
	"sync"

	"github.com/ananrafs/goerranger/goerranger"
)

type workerPool struct {
	maxWorker   int
	queuedTaskC chan func()
	wg          sync.WaitGroup
}

func New(opts goerranger.Options) goerranger.MegaZord {
	wp := &workerPool{
		queuedTaskC: make(chan func(), opts.Count),
		maxWorker:   opts.Count,
	}

	wp.run()

	return wp
}

func (wp *workerPool) Hit(act func()) {
	wp.wg.Add(1)
	go func() {
		wp.addTask(act)
	}()
}

func (wp *workerPool) run() {
	for i := 0; i < wp.maxWorker; i++ {
		wID := i + 1

		go func(workerID int) {
			for hit := range wp.queuedTaskC {
				hit()
				wp.wg.Done()
			}
		}(wID)
	}
}

func (wp *workerPool) GetDisposer() (disposer goerranger.Disposer) {
	return func() {
		wp.wg.Wait()
		close(wp.queuedTaskC)
	}
}

func (wp *workerPool) addTask(task func()) {
	wp.queuedTaskC <- task
}
