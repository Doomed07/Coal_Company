package items

type ItemsInfo struct {
	Pick        Pick
	Ventilation Ventilation
	Minecarts   Minecarts
}

var StaticItemsInfo = ItemsInfo{
	Pick:        *NewPick(),
	Ventilation: *NewVentilation(),
	Minecarts:   *NewMinecarts(),
}
