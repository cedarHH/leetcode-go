package main

import (
	__monostack "leetcode/10_monostack"
	"leetcode/1_array"
	"leetcode/2_linkedList"
	"leetcode/3_hashmap"
	"leetcode/4_strings"
	"leetcode/7_backtracking"
	"leetcode/8_greedy"
	__dynamicProgramming "leetcode/9_dynamicProgramming"
	"leetcode/assessment"
)

func main() {
	arrayTest()
	// linkedListTest()
	// hashmapTest()
	// stringTest()
	// stackQueueTest()
	// binaryTreeTest()
	// backtrackingTest()
	// greedyTest()
	// dynamicProgrammingTest()
	// monostack()
	// graphTest()
	// assessmentTest()
}

func assessmentTest() {
	// assessment.Echo()
	assessment.FindMode()
	assessment.ThousandSeparator()
	assessment.GetNoZeroIntegers()
	assessment.MinAreaRect()
	assessment.CoinChange()
	assessment.BitSetTest()
}

func graphTest() {

}

func monostack() {
	__monostack.Trap()
}

func dynamicProgrammingTest() {
	__dynamicProgramming.MinDistance()
	__dynamicProgramming.Dji()
	__dynamicProgramming.Trap()
}

func greedyTest() {
	__greedy.MaxProfit()
}

func backtrackingTest() {
	__backtracking.SolveNQueens()
}

func binaryTreeTest() {

}

func stackQueueTest() {

}

func stringTest() {
	__strings.ReverseString()
	__strings.ReverseStr()
	__strings.ReplaceNumber()
	__strings.ReverseWords()
	__strings.RightRotation()
	__strings.StrStr()
	__strings.RepeatedSubstringPattern()
}

func hashmapTest() {
	__hashmap.IsAnagram()
	__hashmap.Intersection()
	__hashmap.IsHappy()
	__hashmap.TwoSum()
	__hashmap.FourSumCount()
	__hashmap.CanConstruct()
	__hashmap.ThreeSum()
	__hashmap.FourSum()
}

func linkedListTest() {
	__linkedList.RemoveElements()
	__linkedList.LinkedList()
	__linkedList.ReverseList()
	__linkedList.SwapPairs()
	__linkedList.RemoveNthFromEnd()
	__linkedList.GetIntersectionNode()
	__linkedList.DetectCycle()
}

func arrayTest() {
	__array.Search()
	__array.RemoveElement()
	__array.SortedSquares()
	__array.MinSubArrayLen()
	__array.GenerateMatrix()
	__array.RangeSum()
}
