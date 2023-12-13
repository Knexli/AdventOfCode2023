package main

import "common"

func main() {
	lines := common.ReadInputFile()
	sum := 0
	sum2 := 0
	for _, line := range lines {
		diffTree := createDiffTree(line)
		diffTree.expandDiffTree()
		sum += diffTree.getSumOfNextValues()
		diffTree.preExpandDiffTree()
		sum2 += diffTree.getSumOfFirstValues()
	}
	println(sum)
	println(sum2)
}

type diffTree struct {
	diffs    [][]int
	maxWidth int
}

func (dt *diffTree) expandDiffTree() {
	for i := len(dt.diffs) - 1; i >= 0; i-- {
		if isNotZeros(dt.diffs[i]) {
			lastVal := dt.diffs[i][len(dt.diffs[i])-1]
			diffVal := dt.diffs[i+1][len(dt.diffs[i+1])-1]
			dt.diffs[i] = append(dt.diffs[i], lastVal+diffVal)
		} else {
			dt.diffs[i] = append(dt.diffs[i], 0)
		}
	}
}

func (dt *diffTree) preExpandDiffTree() {
	for i := len(dt.diffs) - 1; i >= 0; i-- {
		if isNotZeros(dt.diffs[i]) {
			firstVal := dt.diffs[i][0]
			diffVal := dt.diffs[i+1][0]
			dt.diffs[i] = append([]int{firstVal - diffVal}, dt.diffs[i]...)
		} else {
			dt.diffs[i] = append([]int{0}, dt.diffs[i]...)
		}
	}
}

func (dt *diffTree) getSumOfNextValues() int {
	return dt.diffs[0][len(dt.diffs[0])-1]
}

func (dt *diffTree) getSumOfFirstValues() int {
	return dt.diffs[0][0]
}

func createDiffTree(line string) diffTree {
	nums := common.ConvertToNumArray(line)
	dt := diffTree{
		maxWidth: len(nums),
		diffs:    make([][]int, 0),
	}
	dt.diffs = append(dt.diffs, nums)
	for isNotZeros(nums) {
		nums = getDiffArray(nums)
		dt.diffs = append(dt.diffs, nums)
	}
	return dt
}

func getDiffArray(nums []int) []int {
	result := make([]int, len(nums)-1)
	for i := 0; i < len(result); i++ {
		result[i] = nums[i+1] - nums[i]
	}
	return result
}

func isNotZeros(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return true
		}
	}
	return false
}
