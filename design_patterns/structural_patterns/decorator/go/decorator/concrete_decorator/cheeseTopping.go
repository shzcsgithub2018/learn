package concrete_decorator

import (
	"github.com/shzgithub2018/learn/design_patterns/structural_patterns/decorator/go"
)

type CheeseTopping struct {
	Pizza _go.IPizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}
