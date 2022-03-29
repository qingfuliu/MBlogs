package generate

import (
	"sync"
	"time"
)

type stopWatch struct {
	time int64
}

func (s *stopWatch) nowMicros(nowMicros int64) int64 {
	return nowMicros - s.time
}
func (s *stopWatch) getMicro() int64 {
	return time.Now().UnixMicro() - s.time
}

type rateLimiter struct {
	mutex sync.Mutex
	stopWatch
}

type smoothRateLimiter struct {
	rateLimiter
	storedToken          int64
	maxStoredToken       int64
	stableIntervalMicros int64
	nextFreeMicros       int64
}

type SmoothBursty struct {
	smoothRateLimiter
	maxBurstSeconds int64
}

func (s *SmoothBursty) Acquire() int64 {
	return s.acquire(1)
}

func (s *SmoothBursty) SetRate(countPerSecond int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	now := time.Now().UnixMicro()
	s.reSync(now)
	intervalMicros := countPerSecond * int64(time.Microsecond) / countPerSecond
	s.stableIntervalMicros = intervalMicros

	prevMaxToken := s.maxStoredToken
	s.maxStoredToken = s.maxBurstSeconds * countPerSecond
	s.storedToken = s.storedToken * s.maxStoredToken / prevMaxToken
}

func (s *SmoothBursty) acquire(count int64) int64 {
	now := time.Now().UnixMicro()
	momentAvailable := s.reserveEarliestAvailable(count, now)
	waitTime := momentAvailable - now
	time.Sleep(time.Duration(waitTime))
	return waitTime / int64(time.Microsecond)
}

func (s *SmoothBursty) reserveEarliestAvailable(count, nowMicros int64) int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.reSync(nowMicros)
	returnValue := s.nextFreeMicros

	availableToken := min(count, s.storedToken)
	deficientToken := count - availableToken
	deficientMicros := deficientToken * s.stableIntervalMicros

	s.nextFreeMicros += deficientMicros
	s.storedToken -= count
	return returnValue
}
func (s *SmoothBursty) reSync(nowMicros int64) {
	nowMicros = s.nowMicros(nowMicros)
	if nowMicros > s.nextFreeMicros {
		var count int64 = (nowMicros - s.nextFreeMicros) / s.stableIntervalMicros
		s.storedToken = min(s.storedToken+count, s.maxStoredToken)
		s.nextFreeMicros = nowMicros
	}
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func GetSmoothBursty(countPreSecond int64) (smoothBursty *SmoothBursty) {
	now := time.Now().UnixMicro()
	smoothBursty = &SmoothBursty{
		maxBurstSeconds: 1,
		smoothRateLimiter: smoothRateLimiter{
			storedToken:          0,
			maxStoredToken:       countPreSecond,
			stableIntervalMicros: int64(time.Microsecond) / countPreSecond,
			nextFreeMicros:       now,
			rateLimiter: rateLimiter{
				stopWatch: stopWatch{
					time: now,
				},
				mutex: sync.Mutex{},
			},
		},
	}
	return
}
