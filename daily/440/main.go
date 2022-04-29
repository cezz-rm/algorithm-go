package main

import "fmt"

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func getCount(prefix, n int) (count int) {
	curr := prefix
	next := prefix + 1
	for curr <= n {
		count += min(n+1, next) - curr
		curr *= 10
		next *= 10
	}
	return
}

//func getSteps(cur, n int) (steps int) {
//	first, last := cur, cur
//	for first <= n {
//		steps += min(last, n) - first + 1
//		first *= 10
//		last = last*10 + 9
//	}
//	return
//}

func findKthNumber(n, k int) int {
	if n/10 < 1 {
		return k
	}
	p := 1      // 作为第一个指针, 指向当前所在位置, 当p=k时, 也就是到了排位第k的数
	prefix := 1 // 前缀
	for p < k {
		count := getCount(prefix, n) // 获取当前前缀下所有子节点的数量
		if p+count > k {             // 第k个数在当前前缀下
			prefix *= 10
			p++ // 把指针指向第一个子节点的位置, 比如11乘10后变成110, 指针从11指向了110
		} else { // 第k个数不在当前前缀下
			prefix++
			p += count // 把指针指向下一前缀的起点
		}
	}
	return prefix
}

func main() {
	//fmt.Println(getCount(3, 13))
	fmt.Println(findKthNumber(10, 3))
}
