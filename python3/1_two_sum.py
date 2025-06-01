

import unittest


class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        d = {}
        for i, num in enumerate(nums):
            tmp = target - num
            j = d.get(tmp)
            if j is not None:
                return [j, i]
            d[num] = i
        return []


class SolutionTestCase(unittest.TestCase):

    def test_twoSum(self):
        class Case:
            def __init__(self, nums, target, ans):
                self.nums = nums
                self.target = target
                self.ans = ans
        
        cases = [
            Case([2, 7, 11, 15], 9, [0, 1]),
            Case([3, 2, 4], 6, [1, 2]),
            Case([3, 3], 6, [0, 1])
        ]
        
        s = Solution()
        for case in cases:
            self.assertEqual(s.twoSum(case.nums, case.target).sort(), case.ans.sort())
       

unittest.main()
     