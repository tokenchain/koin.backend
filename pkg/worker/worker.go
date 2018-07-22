package worker

import (
	"flag"
	"log"
	"os"
	"strconv"
)

// Job holds the attributes needed to perform unit of work.
type Job struct {
	Name string
	Run  func() error
}

var (
	DefaultMaxWorkers   = flag.Int("max_workers", 10, "The number of workers to start")
	DefaultMaxQueueSize = flag.Int("max_queue_size", 1000, "The size of job queue")
	DefaultJobQueue     = make(chan Job, *DefaultMaxQueueSize)
)

// NewWorker creates takes a numeric id and a channel w/ worker pool.
func NewWorker(id int, workerPool chan chan Job) Worker {
	return Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
		logger:     log.New(os.Stderr, "[WORKER "+strconv.Itoa(id)+"] ", 0),
	}
}

// Worker represent the environment where the job will be executed.
type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
	logger     *log.Logger
}

// Add add the job to the default job queue.
// Obviously you can add the job to YOUR custom job queue.
func Add(job Job) {
	DefaultJobQueue <- job
}

// start this is the place where we receive the jobs. The runnable will be executed,
// if there is an error, we log it. We also receive an end-of-worker event,
// in which case we close it.
func (w Worker) start() {
	go func() {
		for {
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				if err := job.Run(); err != nil {
					w.logger.Printf("job %s has an error %s\n", job.Name, err.Error())
				}
			case <-w.quitChan:
				w.logger.Printf("stopping\n")
				return
			}
		}
	}()
}

// stop the worker
func (w Worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}

// NewDispatcher creates, and returns a new Dispatcher object.
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)

	return &Dispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		workerPool: workerPool,
	}
}

type Dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(i+1, d.workerPool)
		worker.start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func() {
				workerJobQueue := <-d.workerPool
				workerJobQueue <- job
			}()
		}
	}
}
