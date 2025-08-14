package main

import (
	. "bufio"
	. "os"
	. "strconv"
)

func main() {
	r := "MCXIDLV"
	f, _ := Open(Args[1])
	s := NewScanner(f)
	for s.Scan() {
		n, _ := Atoi(s.Text())
		o := ""
		a := 1000
		for i := range 4 {
			c := n / a
			x, h := r[i:i+1], r[i+3:i+4]
			if c == 4 {
				o += x + h
			} else if c == 9 {
				o += x + r[i-1:i]
			} else {
				if c >= 5 {
					o += h
				}
				for range c % 5 {
					o += x
				}
			}
			n -= c * a
			a /= 10
		}
		println(o)
	}
}
