package divide

import (
	"fmt"
	"math/big"
)

var Min = big.NewInt(-(1 << 31))
var Max = big.NewInt((1 << 31) - 1)

func divide(dividend int, divisor int) int {
	absDividend := abs(big.NewInt(int64(dividend)))
	absDivisor := abs(big.NewInt(int64(divisor)))

	if absDividend.Cmp(absDivisor) < 0 {
		return 0
	}

	memory := [][2]*big.Int{{big.NewInt(1), absDivisor}}

	out := big.NewInt(0)
	remaining := absDividend

	for remaining.Cmp(absDivisor) >= 0 {
		mem := memory[len(memory)-1]
		memory = append(memory, [2]*big.Int{new(big.Int).Add(mem[0], mem[0]), new(big.Int).Add(mem[1], mem[1])})

		i := len(memory) - 1

		for i > 0 && remaining.Cmp(memory[i][1]) < 0 {
			i--
		}

		out.Add(out, memory[i][0])
		remaining.Sub(remaining, memory[i][1])
	}

	fmt.Printf("out=%d, remaining=%d\n", out, remaining)
	if (dividend < 0 && divisor >= 0) || (dividend >= 0 && divisor < 0) {
		out.Neg(out)
	}

	if out.Cmp(Min) < 0 {
		return int(Min.Int64())
	}

	if out.Cmp(Max) > 0 {
		return int(Max.Int64())
	}

	return int(out.Int64())
}

func abs(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) < 0 {
		return n.Neg(n)
	}

	return n
}
