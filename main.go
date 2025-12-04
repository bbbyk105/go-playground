package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("0で割れないです")
	}
	return a / b, nil
}

func apply(n int, fn func(int) int) int {
		return fn(n)
	}

func double(x int) int {
		return x * 2
}

func main() {
	score := 75

	if score >= 90 {
		fmt.Println("優秀！！")
	} else if score >= 60 {
		fmt.Println("合格")
	} else {
		fmt.Println("不合格")
	}

	fmt.Println("1〜5を表示")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("カウントアップ (3で停止)")
	x := 1
	for {
		fmt.Println(x)
		if x == 3 {
			break
		}
		x++
	}

	// スライス
	nums := []int{10, 20, 30}
	fmt.Println("スライスを回す")
	for index, value := range nums {
		fmt.Println(index, value)
	}

	nums2 := []int{20, 40, 60}
	fmt.Println("sliceを回す")

	// 通常
	for index, value := range nums2 {
		fmt.Println(index, value)
	}

	// index使わない場合
	for _, value := range nums2 {
		fmt.Println(value)
	}

	// 逆スライス
	for i := len(nums2) - 1; i >= 0; i-- {
		fmt.Println(nums2[i])
	}

	// 行列
	matrix2 := [][]int{
		{2, 3},
		{3, 4},
	}

	for i, row2 := range matrix2 {
		for j, value := range row2 {
			fmt.Println(i, j, value)
		}
	}

	// divide 関数の使用
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println(result)
	}

	//チャンネル作成
	ch := make(chan int)

	go func(){
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for v :=range ch {
		fmt.Println(v)
	}

	fmt.Println(apply(10, double))

	
}
