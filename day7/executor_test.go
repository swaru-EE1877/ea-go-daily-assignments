package executor

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	s := Service{}
	expectedCnt := 3
	var actualCnt int32 = 0
	setup := func() {
		atomic.AddInt32(&actualCnt, 1)
	}

	s.Start(expectedCnt, setup)
	assert.Equal(t, expectedCnt, int(actualCnt))
}

func TestRun(t *testing.T) {
	s := Service{}
	expectedCnt := 4
	var actualCnt int32 = 0
	s.Start(3, func() {})
	wg := sync.WaitGroup{}
	job := func() {
		atomic.AddInt32(&actualCnt, 1)
		wg.Done()
	}
	wg.Add(4)
	for i := 0; i < expectedCnt; i++ {
		s.Run(job)
	}
	wg.Wait()
	assert.Equal(t, expectedCnt, int(actualCnt))
}

func TestClose(t *testing.T) {
	s := Service{}
	s.Start(3, func() {})

	s.Close()

	executed := false
	job := func() {
		executed = true
	}
	err := s.Run(job)

	assert.Error(t, err)
	assert.False(t, executed)
}

func TestRunBatch(t *testing.T) {
	s := Service{}
	s.Start(3, func() {})
	expectedCnt := 3
	var actualCnt int32 = 0
	wg := sync.WaitGroup{}
	job1 := func() {
		atomic.AddInt32(&actualCnt, 1)
		wg.Done()
	}
	job2 := func() {
		atomic.AddInt32(&actualCnt, 1)
		wg.Done()
	}
	job3 := func() {
		atomic.AddInt32(&actualCnt, 1)
		wg.Done()
	}
	wg.Add(3)
	s.RunBatch([]func(){
		job1, job2, job3,
	})
	wg.Wait()
	assert.Equal(t, expectedCnt, int(actualCnt))
}
