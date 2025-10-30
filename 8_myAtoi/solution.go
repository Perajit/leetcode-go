package myatoi

import (
	"math/big"
)

var Min = big.NewInt(-(1 << 31))
var Max = big.NewInt((1 << 31) - 1)

var valueLookup = map[byte]*big.Int{
	'0': big.NewInt(0),
	'1': big.NewInt(1),
	'2': big.NewInt(2),
	'3': big.NewInt(3),
	'4': big.NewInt(4),
	'5': big.NewInt(5),
	'6': big.NewInt(6),
	'7': big.NewInt(7),
	'8': big.NewInt(8),
	'9': big.NewInt(9),
}

func myAtoi(s string) int {
	L := len(s)

	from := -1
	to := -1
	var sign byte

	for i := 0; i < L; i++ {
		c := s[i]
		_, ok := valueLookup[c]

		if from < 0 {
			if ok {
				from = i
				to = i
				continue
			} else if sign == 0 {
				if c == '+' || c == '-' {
					sign = c
					continue
				} else if c == ' ' {
					continue
				}
			}

			return 0
		} else {
			if !ok {
				break
			}

			to = i
		}
	}

	if from < 0 {
		return 0
	}

	out := big.NewInt(0)
	m := big.NewInt(1)
	ten := big.NewInt(10)

	for i := to; i >= from; i-- {
		d := valueLookup[s[i]]

		// v = d*m
		v := new(big.Int).Mul(d, m)

		// out += v
		out.Add(out, v)

		// m *= 10
		m.Mul(m, ten)
	}

	if sign == '-' {
		// out = -out
		out.Neg(out)
	}

	// if out < Min
	if out.Cmp(Min) < 0 {
		// out = Min
		out.Set(Min)
	}

	// if out > Max
	if out.Cmp(Max) > 0 {
		// out = Max
		out.Set(Max)
	}

	return int(out.Int64())
}
