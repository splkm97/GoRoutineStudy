package main

import (
	"fmt"
)

/*
ch: 통신용 채널
n: 총 출력 비트수
index: 현재 기준이 되는 비트
prev: 이전에 출력한 그레이 코드
*/
func grayGen(ch chan []int, n int, index int, code []int, rev bool) {
	if index == n {
		ch <- append(code[:0:0], code...)
		return
	} else {
		if rev {
			code[index] = 1
		} else {
			code[index] = 0
		}
		grayGen(ch, n, index+1, code, false)
		if rev {
			code[index] = 0
		} else {
			code[index] = 1
		}
		grayGen(ch, n, index+1, code, true)
	}
}

func GrayCodes(n int) <-chan []int {
	grayStream := make(chan []int)
	go func() {
		defer close(grayStream)
		grayGen(grayStream, n, 0, make([]int, n, n), false)
	}()
	return grayStream
}

func main() {
	for c := range GrayCodes(5) {
		fmt.Println(c)
	}
}
