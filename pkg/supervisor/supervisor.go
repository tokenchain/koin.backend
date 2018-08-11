package supervisor

import (
	"time"
	"log"
	"os"
)

type Run func() error

type Supervisor struct {
	Attempt int
	Time    time.Duration
}

type Then struct {
	execute func(bool)
}

func (t *Then) Then(b func(b bool)) {
	t.execute = b
}

func New(howManyAttempt int, timeBetweenTry time.Duration) Supervisor {
	return Supervisor{howManyAttempt, timeBetweenTry}
}

func (s Supervisor) GoAsync(id string, runnable Run) *Then {
	t := &Then{}
	go func() {
		t.execute(s.GoSync(id, runnable))
	}()
	return t
}

func (s Supervisor) GoSync(id string, runnable Run) bool {
	l := log.New(os.Stdout, "[SUPERVISOR] "+id+": ", 0)
	at := 0

	err := runnable()
	if err == nil {
		return true
	} else {
		for range time.Tick(s.Time) {
			l.Println("New attempt...")
			err := runnable()
			if err == nil {
				return true
			} else {
				l.Println("Attempt fail...")
				at++
				if at > s.Attempt {
					l.Println("Maximum attempt reach, exited supervisor.")
					return false
				}
			}
		}
	}
	return false
}
