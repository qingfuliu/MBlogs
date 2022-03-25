package generate

import (
	"sync"
	"time"
)

const (
	workerId        int64 = 31 //机器id
	workerShift     int8  = 12 //机器id偏移量
	datacenterId    int64 = 31 //机房id
	datacenterShift int8  = 17 //机房id偏移量
	timeStampShift  int8  = 22 ///时间戳偏移量
	timeStampStart  int64 = 1648212702157
	timeStampMask   int64 = (1 << 41) - 1
	sequenceMask    int64 = (1 << 12) - 1
)

var (
	timeNowStamp  int64
	sequence      int64
	lastTimeStamp = time.Now().UnixMilli() & timeStampMask
)
var mutex sync.Mutex

func SnowFlakeUID() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	timeNowStamp = time.Now().UnixMilli() & timeStampMask
	if timeNowStamp == lastTimeStamp {
		sequence = (sequence + 1) & sequenceMask
		if sequence == 0 {
			timeNowStamp = nextMilli()
		}
	} else {
		sequence = 0
	}
	lastTimeStamp = timeNowStamp
	return ((timeNowStamp - timeStampStart) << timeStampShift) |
		(workerId << workerShift) |
		(datacenterId << datacenterShift) |
		(sequence)
}

func nextMilli() int64 {
	for timeNowStamp == time.Now().UnixMilli() {
	}
	return time.Now().UnixMilli()
}
