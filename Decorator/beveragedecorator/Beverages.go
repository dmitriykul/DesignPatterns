package beveragedecorator

type houseBlend struct {
	description string
}

func NewHouseBlend() Beverage {
	return &houseBlend{
		description: "House Blend Coffe",
	}
}

func (h *houseBlend) GetDescription() string {
	return h.description
}

func (h *houseBlend) Cost() float32 {
	return 0.89
}

type espresso struct {
	description string
}

func NewEspresso() Beverage {
	return &espresso{
		description: "Espresso",
	}
}

func (e *espresso) GetDescription() string {
	return e.description
}

func (e *espresso) Cost() float32 {
	return 1.99
}