package main

import (
	"github.com/samber/lo"
	plo "github.com/samber/lo/parallel"
	"time"
)

func main() {
	async := lo.Async[int](func() int {
		time.Sleep(2 * time.Second)
		return 1
	})
	sli := []int{1, 2, 3}
	plo.ForEach(lo.Map(sli, func(i int, idx int) int {
		return i * 2
	}), func(i int, idx int) {
		println(i)
	})
	println(<-async)
}
