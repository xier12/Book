package tools

import (
	"errors"
	"sync"
	"time"
)

var Snow *SnowflakeIDGenerator

const (
	workerIDBits      = 5
	sequenceBits      = 12
	maxWorkerID       = -1 ^ (-1 << workerIDBits)
	maxSequence       = -1 ^ (-1 << sequenceBits)
	timeShift         = workerIDBits + sequenceBits
	workerIDShift     = sequenceBits
	maxBackwardMillis = 5
)

type SnowflakeIDGenerator struct {
	workerID      int64
	lastTimestamp int64
	sequence      int64
	lock          sync.Mutex
}

func NewSnowflakeIDGenerator(workerID int64) (*SnowflakeIDGenerator, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker ID out of range")
	}
	return &SnowflakeIDGenerator{
		workerID: workerID,
	}, nil
}

func (g *SnowflakeIDGenerator) GenerateID() (int64, error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	if timestamp < g.lastTimestamp {
		return 0, errors.New("clock moved backwards")
	}

	if timestamp == g.lastTimestamp {
		g.sequence = (g.sequence + 1) & maxSequence
		if g.sequence == 0 {
			timestamp = g.waitNextMillis(g.lastTimestamp)
		}
	} else {
		g.sequence = 0
	}

	g.lastTimestamp = timestamp

	id := (timestamp << timeShift) | (g.workerID << workerIDShift) | g.sequence
	return id, nil
}

func (g *SnowflakeIDGenerator) waitNextMillis(lastTimestamp int64) int64 {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	for timestamp <= lastTimestamp {
		time.Sleep(time.Millisecond)
		timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	}
	return timestamp
}
