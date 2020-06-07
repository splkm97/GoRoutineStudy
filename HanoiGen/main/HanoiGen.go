package main

import "fmt"

func hanoiGen(ch chan [2]string, n int, from, to, by string) {
	if n == 1 {
		ans := convToArray(from, to)
		ch <- ans
		return
	} else {
		hanoiGen(ch, n-1, from, by, to)
		ans := convToArray(from, to)
		ch <- ans
		hanoiGen(ch, n-1, by, to, from)
	}
}

func convToArray(from string, to string) [2]string {
	ans := [2]string{"", ""}
	ans[0] = from
	ans[1] = to
	return ans
}

func HanoiTowers(n int, from, to, by string) <-chan [2]string {
	hanoiStream := make(chan [2]string)
	go func() {
		defer close(hanoiStream)
		hanoiGen(hanoiStream, n, from, to, by)
	}()
	return hanoiStream
}

func main() {
	for c := range HanoiTowers(3, "A", "B", "temp") {
		fmt.Println(c[0], " -> ", c[1])
	}
}
