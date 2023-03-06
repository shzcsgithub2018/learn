package main

import "log"

func main() {
	var a map[int64]bool
	a = nil

	log.Println(a[100])

}

