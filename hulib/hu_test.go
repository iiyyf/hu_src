package hulib

import (
	"fmt"
	"testing"
	"time"
)

var (
	errCount int = 0
	okCount  int = 0
	tested       = map[int32]bool{}
)

func print_cards(cards []int32) {
	for i := 0; i < 9; i++ {
		fmt.Printf("%d,", cards[i])
	}
	fmt.Printf("\n")

	for i := 9; i < 18; i++ {
		fmt.Printf("%d,", cards[i])
	}
	fmt.Printf("\n")

	for i := 18; i < 27; i++ {
		fmt.Printf("%d,", cards[i])
	}
	fmt.Printf("\n")

	for i := 27; i < 34; i++ {
		fmt.Printf("%d,", cards[i])
	}
	fmt.Printf("\n")
}

func check_hu(cards []int32, max int) {
	for i := 0; i < max; i++ {
		if cards[i] > 4 {
			return
		}
	}

	num := int32(0)
	for i := 0; i < 9; i++ {
		num = num*10 + cards[i]
	}

	_, ok := tested[num]
	if ok {
		return
	}

	tested[num] = true

	for i := 0; i < max; i++ {
		if !CheckHu(cards, 34) {
			errCount++
			fmt.Printf("测试失败 i=%d\n", i)
			print_cards(cards)
		} else {
			okCount++
		}
	}
}

func gen_auto_table_sub(cards []int32, level int) {
	for i := 0; i < 32; i++ {
		index := -1
		if i <= 17 {
			cards[i] += 3
		} else if i <= 24 {
			index = i - 18
		} else {
			index = i - 16
		}

		if index >= 0 {
			cards[index] += 1
			cards[index+1] += 1
			cards[index+2] += 1
		}

		if level == 4 {
			check_hu(cards, 18)
		} else {
			gen_auto_table_sub(cards, level+1)
		}

		if i <= 17 {
			cards[i] -= 3
		} else {
			cards[index] -= 1
			cards[index+1] -= 1
			cards[index+2] -= 1
		}
	}
}

func test_two_color() {
	fmt.Println("testing two colors:")
	cards := []int32{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
	}

	for i := 0; i < 18; i++ {
		cards[i] = 2
		fmt.Printf("eye %d\n", i+1)
		gen_auto_table_sub(cards, 1)
		cards[i] = 0
	}
}

func test_time(count int) {
	cards := []int32{
		0, 0, 0, 0, 0, 1, 0, 0, 0,
		1, 1, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 1, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 0, 0, 4, 4,
	}

	print_cards(cards)
	start := time.Now().Unix()
	for i := 0; i < count; i++ {
		CheckHu(cards, 34)
	}
	print_cards(cards)
	fmt.Println("count=", count, "use time=", time.Now().Unix()-start)
}

func oneok() {
	cards := []int32{
		0, 0, 0, 2, 1, 2, 4, 3, 2,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
	}
	if CheckHu(cards, 34) == false {
		print_cards(cards)
		fmt.Printf("一组数据测试失败 ...\n\n")
	} else {
		fmt.Printf("一组数据测试成功 ...\n\n")
	}
}

func TestCheckHu(t *testing.T) {
	oneok()

	start := time.Now().Unix()
	fmt.Println("开始测试...")
	test_time(100000000)
	test_two_color()
	fmt.Println("测试结束...")
	fmt.Printf("失败次数:%d，成功次数:%d\n", errCount, okCount)
	fmt.Println("测试总次数=", okCount+errCount, "用时", time.Now().Unix()-start)
}
