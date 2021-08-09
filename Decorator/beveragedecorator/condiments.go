package beveragedecorator

type milk struct {
	beverage Beverage
}

func NewMilk(beverage Beverage) CondimentDecorator {
	return &milk{
		beverage: beverage,
	}
}

func (m *milk) GetDescription() string {
	return m.beverage.GetDescription() + ", Milk"
}

func (m *milk) Cost() float32 {
	return .10 + m.beverage.Cost()
}

type mocha struct {
	beverage Beverage
}

func NewMocha(beverage Beverage) CondimentDecorator {
	return &mocha{
		beverage: beverage,
	}
}

func (m *mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", Mocha"
}

func (m *mocha) Cost() float32 {
	return .20 + m.beverage.Cost()
}