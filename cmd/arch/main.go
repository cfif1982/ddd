package main

import "ddd/internal"

func main() {
	if err := internal.Run(); err != nil {
		panic(err)
	}
}
