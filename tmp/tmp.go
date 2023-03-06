package main

import "fmt"

type Message string

func NewMessage() Message { return Message("Hello Wire!") }

type Channel struct{ Message Message }

func NewChannel(m Message) Channel { return Channel{Message: m} }

func (c Channel) GetMsg() Message { return c.Message }

type FF struct {
	Channel Channel
}

func NewFF(c Channel) FF { return FF{Channel: c} }

func (f FF) GetChannel() Channel { return f.Channel }

type BroadCast struct{ f FF }

func NewBroadCast(f FF) BroadCast { return BroadCast{f: f} }

func (b BroadCast) Start() {
	msg := b.f.GetChannel().GetMsg()
	fmt.Println(msg)
}
