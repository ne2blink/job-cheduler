package job

import "sync"

type queue struct {
	jobs []Job
	lock sync.Mutex
}

func (q *queue) Push(job Job) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.jobs = append(q.jobs, job)
}

func (q *queue) Pop() Job {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.jobs) == 0 {
		return nil
	}
	job := q.jobs[0]
	q.jobs = q.jobs[1:]
	return job
}

func (q *queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.jobs)
}
