package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

type Vertex struct {
	X int
	Y int
}

var c, python bool
var j float64 = 42

func main() {
	fmt.Printf("number %d \n", rand.Intn(10))
	fmt.Printf("number %d \n", rand.Intn(time.Now().Nanosecond()))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(time.Now().Nanosecond())
	a, b := swap("World", "Hello")
	fmt.Println(a, b)
	var i int
	java := true
	fmt.Println(i, j, c, python, java)
	k := 42
	f := float64(k)
	fmt.Println(f)
	var MaxInt uint64 = 1<<64 - 1
	fmt.Printf("Type %T Value %v \n", MaxInt, MaxInt)

	l := 0
	for l < 10 {
		l++
		fmt.Println(l)
	}

	os := runtime.GOOS
	fmt.Println(os)

	switch os {
	case "windows":
		fmt.Println("bar")
	default:
		fmt.Println("default")
	}

	//defer fmt.Println("deferred")
	//defer fmt.Println("deferred2")
	fmt.Println("now")
	m := 42
	pointer := &m
	fmt.Println(*pointer)
	fmt.Println(pointer)
	*pointer = 815
	fmt.Println(m)

	n := Vertex{1, 2}
	fmt.Println(n.Y)

	o := Vertex{}
	fmt.Println(o.X)

	var p [2]string
	p[0] = "foo"
	p[1] = "bar"
	fmt.Printf("%v %v \n", p[0], p[1])

	primes := [3]int{2, 3, 5}
	fmt.Println(primes)
	fmt.Println(len(primes))
	for i := 0; i < len(primes); i++ {
		fmt.Println(primes[i])
	}

	var q []int = primes[2:3]
	q[0] = 42
	fmt.Println(q)
	fmt.Println(primes)

	var q1 []int
	fmt.Println(q1, len(q1), cap(q1))

	q2 := make(map[string]int)
	q2["foo"] = 42
	fmt.Println(q2)
	fmt.Println(q2["foo"])

	_, ok := q2["foo"]
	fmt.Println(ok)
	delete(q2, "foo")
	_, ok = q2["foo"]
	fmt.Println(ok)

	x := 0
	y := 1
	for i := 0; i < 50; i++ {
		x, y = fibon(x, y)
		fmt.Println(x)
	}
}

func swap(x, y string) (string, string) {
	return y, x
}

func fibon(x, y int) (int, int) {
	return y, x + y
}
