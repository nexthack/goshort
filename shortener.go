package goshort

import (
	"fmt"
	"strings"
)

const (
	space = "23456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ-_"
	base  = len(space)
)

//Encode returns a string by encoding the id over a 51 characters space
func Encode(id int) string {
	var short []byte
	for id > 0 {
		i := id % base
		short = append(short, byte(space[i]))
		id = id / base
	}
	for i, j := 0, len(short)-1; i < j; i, j = i+1, j-1 {
		short[i], short[j] = short[j], short[i]
	}
	return string(short)
}

//Decode will decode the string and return the id
//The input string should be a valid one with only characters in the space
func Decode(short string) (int, error) {
	if len(short) != len([]rune(short)) {
		return 0, fmt.Errorf("goshort:Input string contains invalid encoded bytes")
	}
	id := 0
	for i := 0; i < len(short); i++ {
		id = id*base + strings.IndexByte(space, short[i])
	}
	return id, nil
}
