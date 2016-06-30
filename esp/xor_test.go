package esp

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"testing"
	"time"
)

func xorEval(nn *NNet) float64 {
	xor := [][]float64{
		[]float64{1.0, 1.0, 0.0},
		[]float64{1.0, 0.0, 1.0},
		[]float64{0.0, 1.0, 1.0},
		[]float64{0.0, 0.0, 0.0},
	}
	score := 0.0
	for _, op := range xor {
		sol := op[2]
		output := nn.Update(op[:2])
		score += math.Pow((sol - output[0]), 2)
	}
	return score
}

func TestXor(testing.T) {
	s := time.Now().UnixNano()
	rand.Seed(s)
	param, err := NewESPParam("xortest.esp")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Seed: %d\n", s)
	param.Show()

	e := New(param)
	e.Run(xorEval)
	nn := e.BestNNet()
	ans := nn.Update([]float64{1.0, 1.0})
	fmt.Printf("1 xor 1 = %f\n", ans[0])
	ans = nn.Update([]float64{1.0, 0.0})
	fmt.Printf("1 xor 0 = %f\n", ans[0])
	ans = nn.Update([]float64{0.0, 1.0})
	fmt.Printf("0 xor 1 = %f\n", ans[0])
	ans = nn.Update([]float64{0.0, 0.0})
	fmt.Printf("0 xor 0 = %f\n", ans[0])
}
