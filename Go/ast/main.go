package main

import(
	"errors"
)

func main(){
	errors.New("改造")
	errors.New("--代码--")
	errors.New(Foo("改造代码"))
}

func Ts(){
	errors.New("zzz")
}

func Foo(s string){}