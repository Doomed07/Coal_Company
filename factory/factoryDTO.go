package factory

import (
	"Coal_company/factory/items"
	"Coal_company/factory/miners"
	"time"
)

type FactoryInfoDTO struct {
	Items     items.ItemsModul
	Miners    miners.Miners
	StartTime time.Time
	WorkTime  time.Duration
	Balance   int
}

type MinersInfoDTO struct {
	Junior []*miners.JuniorMiner
	Middle []*miners.MiddleMiner
	Senior []*miners.SeniorMiner
}
