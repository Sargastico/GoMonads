package main

import (
	"GoEffectSystems/monad"
	"log"
)

func addOne(v interface{}) monad.Maybe {

	if v.(int) > 2 {
		
		return monad.Return(v.(int) + 1)

	}

	return monad.Maybe{}

}

func double(v interface{}) monad.Maybe {

	return monad.Return(v.(int) * 2)

}
func minusOne(v interface{}) monad.Maybe {

	return monad.Return(v.(int) - 1)

}

func main() {

	m := monad.Return(3)
	n := m.Bind(addOne).Bind(double).Bind(minusOne)

	log.Println(n)

}
