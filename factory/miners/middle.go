package miners

import (
	"context"
	"sync"
	"time"
)

var mtx2 sync.RWMutex

type MiddleMiner struct {
	Name     string
	Cost     int
	Stamina  int
	Earn     int
	RestTime int
}

func NewMiddleMiner() *MiddleMiner {
	return &MiddleMiner{
		Name:     "Middle Miner",
		Cost:     50,
		Stamina:  45,
		Earn:     3,
		RestTime: 2,
	}
}

func MiddleMining(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan int,
	middle *MiddleMiner) {
	defer wg.Done()

	totalStamina := middle.Stamina

	for i := 1; i <= totalStamina; i++ {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(middle.RestTime) * time.Second):
		}

		select {
		case <-ctx.Done():
			return
		case transferPoint <- middle.Earn:
			mtx2.Lock()
			middle.Stamina--
			mtx2.Unlock()
		}
	}
}

func (m *MiddleMiner) Buy(ctx context.Context) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go MiddleMining(ctx, wg, coalTransferPoint, m)

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}

func (m *MiddleMiner) Info() int {
	return m.Cost
}

func (m *MiddleMiner) AddTo(mn *Miners) {
	mn.Middle = append(mn.Middle, m)
}
