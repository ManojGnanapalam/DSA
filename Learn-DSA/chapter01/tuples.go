package chapter01

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func powerSeriesN(a int) (s int, c int) {
	s = a * a
	c = a * a * a
	return
}

func TuplesDSA() {
	fmt.Println(powerSeries(4))
	fmt.Println(powerSeriesN(3))

}
