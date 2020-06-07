package main

import "fmt"

/*
ch: 통신용 채널
res: 결과 임시 저장
r: 아직 덜 뽑은 개수
a: 주어진 배열
n: 총 뽑을 개수
visit: 뽑은 수
target: 현재 확인하려는 배열의 인덱스
*/
func combGen(ch chan []string, res []string, r int, a []string, n int, visit int, target int) {
	if r == 0 {
		ch <- append(res[:0:0], res...)
		return
	} else if visit == n {
		return
	} else if target == len(a) {
		return
	} else {
		// 본인 안뽑음
		combGen(ch, res, r, a, n, visit, target+1)
		// 본인 뽑음
		r--
		res = append(res, a[target])
		combGen(ch, res, r, a, n, visit+1, target+1)
	}
	return
}

func Combinations(arr []string, m int) <-chan []string {
	combStream := make(chan []string)
	go func() {
		defer close(combStream)
		combGen(combStream, []string{}, m, arr, m, 0, 0)
	}()
	return combStream
}

func main() {
	for c := range Combinations([]string{"사과", "배", "포도", "복숭아"}, 2) {
		fmt.Println(c)
	}
}
