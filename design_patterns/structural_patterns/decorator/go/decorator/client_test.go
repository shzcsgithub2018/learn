package decorator

import (
	"github.com/shzgithub2018/learn/design_patterns/structural_patterns/decorator/go"
	"github.com/shzgithub2018/learn/design_patterns/structural_patterns/decorator/go/concrete_component"
	"github.com/shzgithub2018/learn/design_patterns/structural_patterns/decorator/go/decorator/concrete_decorator"
	"testing"
)

func TestClient(t *testing.T) {
	var pizza _go.IPizza
	pizza = &concrete_component.VeggieMania{}

	//Add cheese topping
	pizza = &concrete_decorator.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizza = &concrete_decorator.TomatoTopping{
		Pizza: pizza,
	}

	t.Log("Price of veggeMania with tomato and cheese topping is", pizza.GetPrice())
}
