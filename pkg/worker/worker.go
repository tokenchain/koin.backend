package worker

import (
	"os"
	"log"
	"github.com/koinkoin-io/koinkoin.backend/pkg/util"
)

var (
	MaxWorker = util.GetEnvOrDefaultInt("MAX_WORKERS", 10)
	MaxQueue  = util.GetEnvOrDefaultInt("MAX_QUEUE", 1000)
)

var workerLog = log.New(os.Stderr, "WORKER: ", 0)

// Job represents the job to be run
type Job struct {
	Name string
	Runnable func() error
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool  chan chan Job
	JobChannel  chan Job
	quit    	chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// PushJob send to the channel a new job.
func PushJob(job Job) {
	JobQueue <- job
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it.
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				if err := job.Runnable(); err != nil {
					workerLog.Printf("Error on a job " + job.Name)
				}
			case <-w.quit:
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}