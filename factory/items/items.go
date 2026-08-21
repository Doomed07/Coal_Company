package items

type Pick struct {
	Name    string
	Cost    int
	Numbers int
}

func NewPick() *Pick {
	return &Pick{
		Name:    "Pick",
		Cost:    3000,
		Numbers: 0,
	}
}

func (p *Pick) Buy() {
	p.Numbers++
}

func (p *Pick) InfoCost() int {
	return p.Cost
}

type Ventilation struct {
	Name    string
	Cost    int
	Numbers int
}

func NewVentilation() *Ventilation {
	return &Ventilation{
		Name:    "Ventilation",
		Cost:    15_000,
		Numbers: 0,
	}
}

func (v *Ventilation) Buy() {
	v.Numbers++
}

func (v *Ventilation) InfoCost() int {
	return v.Cost
}

type Minecarts struct {
	Name    string
	Cost    int
	Numbers int
}

func NewMinecarts() *Minecarts {
	return &Minecarts{
		Name:    "Minecart",
		Cost:    50_000,
		Numbers: 0,
	}
}
func (m *Minecarts) Buy() {
	m.Numbers++
}

func (m *Minecarts) InfoCost() int {
	return m.Cost
}
