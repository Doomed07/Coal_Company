package miners

import (
	"context"
	"sync"
	"time"
)

var mtx1 sync.RWMutex

type JuniorMiner struct {
	Name     string
	Cost     int
	Stamina  int
	Earn     int
	RestTime int
}

func NewJuniorMiner() *JuniorMiner {
	return &JuniorMiner{
		Name:     "Junior Miner",
		Cost:     5,
		Stamina:  30,
		Earn:     1,
		RestTime: 3,
	}
}

func JuniorMining(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan int,
	junior *JuniorMiner) {
	defer wg.Done()

	totalStamina := junior.Stamina

	for i := 1; i <= totalStamina; i++ {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(junior.RestTime) * time.Second):
		}

		select {
		case <-ctx.Done():
			return
		case transferPoint <- junior.Earn:
			mtx1.Lock()
			junior.Stamina--
			mtx1.Unlock()
		}
	}
}

func (j *JuniorMiner) Buy(ctx context.Context) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go JuniorMining(ctx, wg, coalTransferPoint, j)

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}

func (j *JuniorMiner) Info() int {
	return j.Cost
}

func (j *JuniorMiner) AddTo(m *Miners) {
	m.Junior = append(m.Junior, j)
}
