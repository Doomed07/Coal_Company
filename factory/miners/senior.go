package miners

import (
	"context"
	"sync"
	"time"
)

var mtx3 sync.RWMutex

type SeniorMiner struct {
	Name     string
	Cost     int
	Stamina  int
	Earn     int
	RestTime int
}

func NewSeniorMiner() *SeniorMiner {
	return &SeniorMiner{
		Name:     "Senior Miner",
		Cost:     450,
		Stamina:  60,
		Earn:     10,
		RestTime: 1,
	}
}

func SeniorMining(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan int,
	senior *SeniorMiner) {
	defer wg.Done()

	totalStamina := senior.Stamina

	for i := 1; i <= totalStamina; i++ {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(senior.RestTime) * time.Second):
		}

		select {
		case <-ctx.Done():
			return
		case transferPoint <- senior.Earn:
			mtx3.Lock()
			senior.Earn += 3
			senior.Stamina--
			mtx3.Unlock()
		}
	}
}

func (s *SeniorMiner) Buy(ctx context.Context) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go SeniorMining(ctx, wg, coalTransferPoint, s)

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}

func (s *SeniorMiner) Info() int {
	return s.Cost
}

func (s *SeniorMiner) AddTo(m *Miners) {
	m.Senior = append(m.Senior, s)
}
