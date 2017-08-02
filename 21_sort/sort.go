package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type IntArray []int

func (a IntArray) Len() int {
	return len(a)
}

func (a IntArray) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a IntArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	ia := IntArray{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(ia)

	sort.Sort(ia)
	fmt.Println(ia)

	rs := rand.NewSource(int64(time.Now().Second()))
	r := rand.New(rs)
	for i := 0; i < 10; i++ {
		ia[i] = r.Int() % 100
	}
	fmt.Println(ia)

	sort.Sort(ia)
	fmt.Println(ia)
}

//type IntArray []int
//func (a *IntArray) Len() int {
//return len(*a)
//}
//func (a *IntArray) Less(i, j int) bool {
//return (*a)[i] < (*a)[j]
//}
//func (a *IntArray) Swap(i, j int) {
//(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
//}
