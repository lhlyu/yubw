package job

type Job interface {
	Do()
}

type Worker struct {
	WorkerPool chan chan Job
	JobQueue   chan Job
}

func NewWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		WorkerPool: workerPool,
		JobQueue:   make(chan Job),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do()
			}
		}
	}()
}

type Dispatcher struct {
	WorkerPool        chan chan Job
	JobChan           chan Job
	MaxWorkerPoolSize int
	Workers           []*Worker
}

// maxWorkerPoolSize - 工作池大小
// bufSize - 任务队列缓存大小
func New(maxWorkerPoolSize, bufSize int) *Dispatcher {
	pool := make(chan chan Job)
	return &Dispatcher{
		WorkerPool:        pool,
		JobChan:           make(chan Job, bufSize),
		MaxWorkerPoolSize: maxWorkerPoolSize,
		Workers:           make([]*Worker, 0),
	}
}

// 阻塞等待，限制goroutine的数量
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkerPoolSize; i++ {
		worker := NewWorker(d.WorkerPool)
		d.Workers = append(d.Workers, worker)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobChan:
			// 从worker池取出一个任务队列
			worker := <-d.WorkerPool
			go func(job Job, worker chan Job) {
				// 将job发送到worker chan
				worker <- job
			}(job, worker)
		}
	}
}

// 非阻塞等待，不限制goroutine的数量
func (d *Dispatcher) NbwRun() {
	for i := 0; i < d.MaxWorkerPoolSize; i++ {
		worker := NewWorker(d.WorkerPool)
		d.Workers = append(d.Workers, worker)
		worker.Start()
	}
	go d.dispatch2()
}

func (d *Dispatcher) dispatch2() {
	for {
		select {
		case job := <-d.JobChan:

			go func(job Job) {
				jobQueue := <-d.WorkerPool
				// 将job发送到worker chan
				jobQueue <- job
			}(job)
		}
	}
}

// TODO 统一关闭
// TODO 静默时关闭worker数量
