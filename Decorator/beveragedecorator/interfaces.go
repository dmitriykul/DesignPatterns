package beveragedecorator

type Beverage interface {
	GetDescription() string
	Cost() float32
}

type CondimentDecorator interface {
	Beverage
}