package main

import (
	. "bufio"
	. "os"
	. "strconv"
)

func main() {
	f, _ := Open(Args[1])
	s, r := NewScanner(f), "MCXIDLV"
	for s.Scan() {
		n, _ := Atoi(s.Text())
		o, a := "", 1000
		for i := range 4 {
			c, x, h := n/a, r[i:i+1], r[i+3:i+4]
			if c == 4 {
				o += x + h
			} else if c == 9 {
				o += x + r[i-1:i]
			} else {
				if c > 4 {
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
