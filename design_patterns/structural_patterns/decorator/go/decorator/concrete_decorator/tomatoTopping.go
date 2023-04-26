package concrete_decorator

import "github.com/shzgithub2018/learn/design_patterns/structural_patterns/decorator/go"

type TomatoTopping struct {
	Pizza _go.IPizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 7
}
