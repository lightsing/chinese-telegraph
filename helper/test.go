package main

import (
	"fmt"
	"github.com/lightsing/chinese-telegraph"
)

func main() {
	converter := chinese_telegraph.NewConverter()
	for i := 0; i < 10000; i++ {
		fmt.Println(converter.GetCharacter(uint(i)))
	}
}
