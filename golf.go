package main

import (
	. "bufio"
	. "math"
	. "os"
	. "strconv"
	t "strings"
)

func main() {
	r := "MCXIDLV"
	f, _ := Open(Args[1])
	s := NewScanner(f)
	for s.Scan() {
		n, _ := Atoi(s.Text())
		o := ""
		for i := range 4 {
			a := int(Pow10(3 - i))
			c := n / a
			x := r[i : i+1]
			h := r[i+3 : i+4]
			if c == 4 {
				o += x + h
			} else if c == 9 {
				o += x + r[i-1:i]
			} else {
				if c >= 5 {
					o += h
				}
				o += t.Repeat(x, c%5)
			}
			n -= c * a
		}
		println(o)
	}
}
