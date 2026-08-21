package factory

import (
	"Coal_company/factory/items"
	"Coal_company/factory/miners"
	"context"
	"errors"
	"sync"
	"time"
)

var ErrorNotEnoughCoal = errors.New("Payment canceled. Don't have enough coal in your balance")

type Factory struct {
	Items     *items.ItemsModul
	Miners    *miners.Miners
	Ctx       context.Context
	Cancel    context.CancelFunc
	mtx       sync.Mutex
	StartTime time.Time
	WorkTime  time.Duration
	Balance   int
}

func NewFactory(items items.ItemsModul, miners miners.Miners) *Factory {
	ctx, cancel := context.WithCancel(context.Background())
	f := &Factory{
		Items:     &items,
		Miners:    &miners,
		Ctx:       ctx,
		Cancel:    cancel,
		StartTime: time.Now(),
		Balance:   0,
	}
	go f.PassiveIncome()
	return f
}

func (f *Factory) Shootinfo() FactoryInfoDTO {
	return FactoryInfoDTO{
		Items:     *f.Items,
		Miners:    f.Miners.MinerInfo(),
		StartTime: f.StartTime,
		WorkTime:  f.WorkTime,
		Balance:   f.Balance,
	}
}

func (f *Factory) ShootMinersInfo(m miners.Miners) MinersInfoDTO {
	return MinersInfoDTO{
		Junior: m.Junior,
		Middle: m.Middle,
		Senior: m.Senior,
	}
}

// Main section

func (f *Factory) PassiveIncome() {
	for {
		select {
		case <-f.Ctx.Done():
			return
		case <-time.After(1 * time.Second):
			f.mtx.Lock()
			f.Balance++
			f.mtx.Unlock()
		}
	}
}

func (f *Factory) Allinfo() FactoryInfoDTO {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	return f.Shootinfo()
}

func (f *Factory) Stop() FactoryInfoDTO {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	f.Cancel()
	f.WorkTime = time.Since(f.StartTime)
	return f.Shootinfo()
}

// Miners' section

func (f *Factory) MinersCostInfo() miners.MinersInfo {
	return miners.DefaultMinersInfo
}

func (f *Factory) BuyMiner(class miners.MinerClasses) error {
	f.mtx.Lock()

	if f.Balance < f.Miners.Info(class) {
		f.mtx.Unlock()
		return ErrorNotEnoughCoal
	}

	f.Balance -= f.Miners.Info(class)
	class.AddTo(f.Miners)
	f.mtx.Unlock()

	go func() {
		for earn := range f.Miners.Buy(f.Ctx, class) {
			f.mtx.Lock()
			f.Balance += earn
			f.mtx.Unlock()
		}
	}()

	return nil
}

func (f *Factory) WorkingMiners() MinersInfoDTO {
	w := f.Miners.WorkingInfo()
	return f.ShootMinersInfo(w)
}

func (f *Factory) AllMiners() MinersInfoDTO {
	w := f.Miners.MinerInfo()
	return f.ShootMinersInfo(w)
}

//Items' section

func (f *Factory) BuyItem(item items.Items) error {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	if f.Balance < f.Items.InfoCost(item) {
		return ErrorNotEnoughCoal
	}

	f.Balance -= f.Items.InfoCost(item)
	f.Items.Buy(item)

	return nil
}

func (f *Factory) InfoItems() items.ItemsInfo {
	return items.StaticItemsInfo
}

func (f *Factory) AllBoughtItems() items.ItemsModul {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	return f.Items.InfoBoughtItems()
}
