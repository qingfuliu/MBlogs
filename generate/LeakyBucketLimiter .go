package generate

import (
	"sync"
	"time"
)

type LeakyBucketLimiter struct {
	lastQueryTime int64
	capacity      int64
	rate          int64
	currentSize   int64
	mutex         sync.Mutex
}

func (l *LeakyBucketLimiter) reSync(timeStamp int64) {
	diff := (timeStamp - l.lastQueryTime) * l.rate
	l.currentSize = min(0, l.currentSize-diff)
	l.lastQueryTime = timeStamp
}

func (l *LeakyBucketLimiter) Acquire(nums int64) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.reSync(time.Now().UnixMicro())
	if nums+l.currentSize > l.capacity {
		return false
	} else {
		l.currentSize += nums
		return true
	}
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
