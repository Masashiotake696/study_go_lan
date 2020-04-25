package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MinInt8)

	fmt.Println(math.Abs(3.14))
	fmt.Println(math.Abs(-3.14))

	fmt.Println(math.Pow(0, 2))
	fmt.Println(math.Pow(1, 10))
	fmt.Println(math.Pow(2, 10))

	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))

	fmt.Println(math.Trunc(1.5))
	fmt.Println(math.Trunc(-1.5))
	fmt.Println(math.Floor(1.5))
	fmt.Println(math.Floor(-1.5))
	fmt.Println(math.Ceil(1.5))
	fmt.Println(math.Ceil(-1.5))

	rand.Seed(256)
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	fmt.Println(rnd.Intn(100))
}
