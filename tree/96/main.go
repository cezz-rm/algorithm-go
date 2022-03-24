package main

import "fmt"

// 96. 不同的二叉树

func numTrees(n int) int {
	num := make([][]int, n+1)
	for i := range num {
		num[i] = make([]int, n+1)
	}
	return count(1, n, &num)
}

func count(low, high int, num *[][]int) (res int) {
	if low > high {
		return 1
	}
	if (*num)[low][high] != 0 {
		return (*num)[low][high]
	}
	for i := low; i <= high; i++ {
		left := count(low, i-1, num)
		right := count(i+1, high, num)
		res += left * right
	}
	(*num)[low][high] = res
	return res
}

// 动态规划
//假设 n 个节点存在二叉排序树的个数是 G (n)，令 f(i) 为以 i 为根的二叉搜索树的个数，则
//G(n) = f(1) + f(2) + f(3) + f(4) + ... + f(n)G(n)=f(1)+f(2)+f(3)+f(4)+...+f(n)

//当 i 为根节点时，其左子树节点个数为 i-1 个，右子树节点为 n-i，则
//f(i) = G(i-1)*G(n-i)f(i)=G(i−1)∗G(n−i)

//综合两个公式可以得到 卡特兰数 公式
//G(n) = G(0)*G(n-1)+G(1)*G(n-2)+...+G(n-1)*G(0)G(n)=G(0)∗G(n−1)+G(1)∗(n−2)+...+G(n−1)∗G(0)

func numTrees2(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
	fmt.Println(numTrees2(3))
}
