class Solution1:
    def jump(self, steps) :
        # write code here
        dp = [100001 for _ in steps]
        dp[0] = 0
        for idx in range(len(dp)):
            for j in range(min(idx+steps[idx]+1, len(steps))- idx):
                dp[idx+j] = min(dp[idx+j], dp[idx]+1)
        return dp[len(steps) - 1]

s = Solution1()
print(s.jump([3, 1, 2, 2, 3, 3, 2, 4, 0, 2, 2, 0]))


class Solution:
    def find(self, arr) :
        # write code here
        self.arr = arr
        self.arrLen = len(arr)
        self.res = set()
        for idx in range(len(arr)):
            self.subFind([arr[idx]], idx, idx)
        return len(self.res)

    def subFind(self, subArr, start, end):
        next = end + 1
        while next < self.arrLen:
            if self.arr[next] >= self.arr[end]:
                nextArr = subArr + [self.arr[next]]
                self.res.add(tuple(nextArr))
                self.subFind(nextArr, start, next)
            next+=1


s = Solution()
print(s.find([4,6,7,7]))
