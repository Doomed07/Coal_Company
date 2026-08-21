package miners

import (
	"context"
	"errors"
)

var ErrorUnknownClass = errors.New("unknown miner class")

type MinerClasses interface {
	Buy(ctx context.Context) <-chan int
	Info() int
	AddTo(*Miners)
}

type Miners struct {
	Junior       []*JuniorMiner
	Middle       []*MiddleMiner
	Senior       []*SeniorMiner
	minerClasses MinerClasses
}

func NewMiners() *Miners {
	return &Miners{}
}

func (m *Miners) ClassByName(name string) (MinerClasses, error) {
	if name == "juniorminer" {
		return NewJuniorMiner(), nil
	} else if name == "middleminer" {
		return NewMiddleMiner(), nil
	} else if name == "seniorminer" {
		return NewSeniorMiner(), nil
	} else {
		return nil, ErrorUnknownClass
	}
}

func (m *Miners) Buy(ctx context.Context, miner MinerClasses) <-chan int {
	return miner.Buy(ctx)
}

func (m *Miners) Info(miner MinerClasses) int {
	return miner.Info()
}

func (m *Miners) AddTo(miner MinerClasses) {
	miner.AddTo(m)
}

func (m *Miners) MinerInfo() Miners {
	mtx1.RLock()
	mtx2.RLock()
	mtx3.RLock()
	defer mtx1.RUnlock()
	defer mtx2.RUnlock()
	defer mtx3.RUnlock()
	return *m
}

func (m *Miners) WorkingInfo() Miners {
	working := Miners{}

	mtx1.RLock()
	for _, miner := range m.Junior {
		if miner.Stamina > 0 {
			working.Junior = append(working.Junior, miner)
		}
	}
	mtx1.RUnlock()

	mtx2.RLock()
	for _, miner := range m.Middle {
		if miner.Stamina > 0 {
			working.Middle = append(working.Middle, miner)
		}
	}
	mtx2.RUnlock()

	mtx3.RLock()
	for _, miner := range m.Senior {
		if miner.Stamina > 0 {
			working.Senior = append(working.Senior, miner)
		}
	}
	mtx3.RUnlock()

	return working
}
