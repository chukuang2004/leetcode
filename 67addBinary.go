package main

import (
	"bytes"
	"fmt"
)

func addBinary(a, b string) string {

	alen := len(a)
	blen := len(b)

	clen := alen
	if blen > clen {
		clen = blen

		for i := 0; i < blen-alen; i++ {
			a = "0" + a
		}
	} else {
		for i := 0; i < alen-blen; i++ {
			b = "0" + b
		}
	}

	return add(a, b, false)
}

func add(a, b string, jinwei bool) string {

	l := len(a)
	tmp := a[l-1]
	if jinwei {
		if tmp != b[l-1] {
			tmp = '0'
			jinwei = true
		} else if tmp == '0' {
			tmp = '1'
			jinwei = false
		} else if tmp == '1' {
			tmp = '1'
			jinwei = true
		}
	} else {
		if tmp != b[l-1] {
			tmp = '1'
			jinwei = false
		} else if tmp == '0' {
			tmp = '0'
			jinwei = false
		} else if tmp == '1' {
			tmp = '0'
			jinwei = true
		}
	}

	ret := bytes.NewBuffer(nil)
	if l == 1 {
		if jinwei {
			ret.WriteByte('1')
		}
		ret.WriteByte(tmp)
		return ret.String()
	}
	ret.WriteByte(tmp)

	return add(a[:l-1], b[:l-1], jinwei) + ret.String()
}

func main() {

	fmt.Println(addBinary("1101", "110"))

}
