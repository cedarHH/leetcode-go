package main

import (
	"context"
	__array "leetcode/01_array"
	__linkedList "leetcode/02_linkedList"
	__hashmap "leetcode/03_hashmap"
	__strings "leetcode/04_strings"
	__stack_queue "leetcode/05_stack_queue"
	__backtracking "leetcode/07_backtracking"
	__greedy "leetcode/08_greedy"
	__dynamicProgramming "leetcode/09_dynamicProgramming"
	__monostack "leetcode/10_monostack"
	"leetcode/assessment"
	"leetcode/concurrency"
	"leetcode/exercises"
	"leetcode/network"
	"sync"
	"time"
)

func main() {
	// arrayTest()
	// linkedListTest()
	// hashmapTest()
	// stringTest()
	// stackQueueTest()
	// binaryTreeTest()
	// backtrackingTest()
	// greedyTest()
	// dynamicProgrammingTest()
	// monostackTest()
	// graphTest()
	// networkTest()
	// concurrencyTest()
	// exercisesTest()
	assessmentTest()
}

func assessmentTest() {
	//assessment.RedisExample()
	//assessment.ShopeeAssessment()
	//assessment.DewuAssessment()
	//assessment.TripAssessment()
	//assessment.SxfAssessment()
	//assessment.JdAssessment()
	//assessment.SfAssessment()
	//assessment.MeituanAssessment()
	//assessment.RedBookAssessment()
	//assessment.XingyeAssessment()
	//assessment.BaiduAssessment()
	//assessment.ShanghaiAssessment()
	//assessment.XiaomiAssessment()
	//assessment.XhsAssessment()
	//assessment.TyyAssessment()
	assessment.DewuAssessment2()
}

func exercisesTest() {
	//exercises.FindMode()
	//exercises.ThousandSeparator()
	//exercises.GetNoZeroIntegers()
	//exercises.MinAreaRect()
	//exercises.CoinChange()
	//exercises.BitSetTest()
	exercises.Misc()
	exercises.AssignTasks()
	exercises.CountRangeSum()
	exercises.NumIslands()
	exercises.CountPrimes()
	exercises.Temp()
}

func networkTest() {
	wg := &sync.WaitGroup{}
	wg.Add(5)
	ctx, cancel := context.WithCancel(context.Background())
	go network.EchoServer(ctx, wg)
	wg.Wait()

	wgc := &sync.WaitGroup{}
	wgc.Add(5)
	network.StartClient(wgc)
	wgc.Wait()
	cancel()
	time.Sleep(1 * time.Second)
}

func concurrencyTest() {
	concurrency.MutexExample()
	concurrency.RWLockExample()
	concurrency.WaitGroupExample()
	concurrency.OnceExample()
	concurrency.CondExample()
	concurrency.SyncMapExample()
	concurrency.ObjectBufferPool()
	concurrency.GoroutinePool()
	concurrency.ChannelExample()
	concurrency.TimerExample()
}

func graphTest() {

}

func monostackTest() {
	__monostack.Trap()
}

func dynamicProgrammingTest() {
	__dynamicProgramming.Fib()
	__dynamicProgramming.ClimbStairs()
	__dynamicProgramming.MinCostClimbingStairs()
	__dynamicProgramming.UniquePaths()
	__dynamicProgramming.UniquePathsWithObstacles()
	__dynamicProgramming.IntegerBreak()
	__dynamicProgramming.NumTrees()
	__dynamicProgramming.CanPartition()
	__dynamicProgramming.LastStoneWeightII()
	__dynamicProgramming.FindTargetSumWays()
	__dynamicProgramming.FindMaxForm()
	__dynamicProgramming.Change()
	__dynamicProgramming.MinDistance()
	__dynamicProgramming.Dji()
	__dynamicProgramming.Trap()
	__dynamicProgramming.CanIWin()
}

func greedyTest() {
	__greedy.MaxProfit()
	__greedy.FindContentChildren()
}

func backtrackingTest() {
	__backtracking.SolveNQueens()
}

func binaryTreeTest() {

}

func stackQueueTest() {
	__stack_queue.MyQueueTest()
	__stack_queue.MyStackTest()
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
