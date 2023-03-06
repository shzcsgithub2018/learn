//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeBroadCast() BroadCast {
	wire.Build(NewBroadCast, NewChannel, NewMessage, NewFF)
	return BroadCast{}
}
