package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	data, _ := ioutil.ReadFile("code2char-src.txt")
	var target [10000]rune
	src := strings.Split(string(data), "\n")
	for _, line := range src {
		e := strings.Split(line, " ")
		index, _ := strconv.ParseUint(e[0], 10, 32)
		_rune, _ := utf8.DecodeRuneInString(e[1])
		target[index] = _rune
	}
	for id := range target {
		fmt.Println(target[id])
	}
	f, _ := os.Create("code2char-rune-src.bin")
	defer f.Close()
	fileTargets := make([]byte, 0, 40000)
	for id := range target {
		runeTarget := make([]byte, 4)
		binary.LittleEndian.PutUint32(runeTarget, uint32(target[id]))
		fileTargets = append(fileTargets, runeTarget...)
		//fmt.Println(fileTargets[id])
	}
	f.Write(fileTargets)
}
