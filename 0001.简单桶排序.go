/*
	简单桶排序(不是完整的桶排序)
	时间复杂度:O(M+N)
		M:桶的个数
		N:待排序的元素的个数
	目的:随机的一组数字,按照从大到小的顺序排序
	说明:这个排序实际上就是统计元素的出现个数,因为统计用的数组是有序的,
		所以最终按照有序数组的顺序排下来,
		在将元素出现的数量表现出来就可以了.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	max := 10
	sliSrc := gen(5, max)
	fmt.Println("原始数据:", sliSrc)
	sliTar := mySort(sliSrc)
	fmt.Println("我的方法:", sliTar)

	fmt.Println("原始数据:", sliSrc)
	sliTar = simpleBucketSort(sliSrc, max)
	fmt.Println("简单桶排序:", sliTar)
}

func gen(length, max int) []int {
	rand.Seed(time.Now().Unix())
	r := make([]int, 0, length)
	for i := 0; i < length; i++ {
		r = append(r, rand.Intn(max))
	}
	return r
}

// 这里是我自己使用的冒泡排序
func mySort(sliSrc []int) []int {
	l := len(sliSrc)
	sliTar := make([]int, l)
	copy(sliTar, sliSrc)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if sliTar[i] < sliTar[j] {
				sliTar[i], sliTar[j] = sliTar[j], sliTar[i]
			}
		}
	}
	return sliTar
}

// 简单桶排序
func simpleBucketSort(sliSrc []int, max int) []int {
	sliBucket := make([]int, max+1)
	for _, v := range sliSrc {
		sliBucket[v]++
	}
	index := 0
	sliTar := make([]int, len(sliSrc))
	copy(sliTar, sliSrc)
	for i := max; i >= 0; i-- {
		for j := sliBucket[i]; j > 0; j-- {
			sliTar[index] = i
			index++
		}
	}
	return sliTar
}
