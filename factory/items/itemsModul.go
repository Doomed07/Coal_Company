package items

import "errors"

var ErrorUnknownItem = errors.New("unknown item")

type Items interface {
	Buy()
	InfoCost() int
}

type ItemsModul struct {
	Items []Items
}

func NewItems() *ItemsModul {
	return &ItemsModul{}
}

func (i *ItemsModul) Buy(item Items) {
	item.Buy()
	i.Items = append(i.Items, item)
}

func (i *ItemsModul) InfoCost(item Items) int {
	return item.InfoCost()
}

func (m *ItemsModul) ItemByName(name string) (Items, error) {
	if name == "pick" {
		return NewPick(), nil
	} else if name == "ventilation" {
		return NewVentilation(), nil
	} else if name == "minecarts" {
		return NewMinecarts(), nil
	} else {
		return nil, ErrorUnknownItem
	}
}

func (i *ItemsModul) InfoBoughtItems() ItemsModul {
	return *i
}
