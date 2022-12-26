package executor

import (
	"errors"
	"sync"
)

type Service struct {
	jobChn   chan func()
	isClosed bool
}

func (s *Service) Start(workerCnt int, setup func()) {
	jobChn := make(chan func(), 2)
	s.jobChn = jobChn
	wg := sync.WaitGroup{}
	wg.Add(workerCnt)
	for i := 0; i < workerCnt; i++ {
		go worker(setup, &wg, jobChn)
	}
	wg.Wait()
}

func worker(setup func(), wg *sync.WaitGroup, jobChn chan func()) {
	setup()
	wg.Done()
	for {
		job, ok := <-jobChn
		if !ok {
			break
		}
		job()
	}
}

func (s Service) Run(job func()) error {
	if s.isClosed {
		return errors.New("closed")
	}
	s.jobChn <- job
	return nil
}

func (s *Service) Close() {
	close(s.jobChn)
	s.isClosed = true
}

func (s Service) RunBatch(jobs []func()) {
	for i := 0; i < len(jobs); i++ {
		s.Run(jobs[i])
	}
}
