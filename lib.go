package chinese_telegraph

import (
	"encoding/binary"
	"github.com/lightsing/chinese-telegraph/bindata"
)

type Converter struct {
	src [10000]rune
}

func NewConverter() *Converter {
	converter := &Converter{}
	data, _ := bindata.Asset("code2char-rune-src.bin")
	for i := range converter.src {
		converter.src[i] = rune(binary.LittleEndian.Uint32(data[i*4 : (i+1)*4]))
	}
	return converter
}

func (converter *Converter) GetCharacter(index uint) rune {
	return converter.src[index]
}

func (converter *Converter) GetCharacterString(index uint) string {
	return string(converter.src[index])
}
