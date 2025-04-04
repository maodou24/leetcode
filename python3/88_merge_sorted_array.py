from typing import List
import unittest

class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        pos = len(nums1) - 1
        i, j = m - 1, n-1
        while i >= 0 and j >= 0:
            if nums1[i] > nums2[j]:
                nums1[pos] = nums1[i]
                i = i-1
            else:
                nums1[pos] = nums2[j]
                j = j-1
            pos = pos-1

        while j >= 0:
            nums1[pos] = nums2[j]
            pos, j = pos-1, j-1

        

class SolutionTestCase(unittest.TestCase):

    def test_merge(self):
        class Case:
            def __init__(self, nums1, m, nums2, n, ans):
                self.nums1 = nums1
                self.m = m
                self.nums2 = nums2
                self.n = n
                self.ans = ans
        
        cases = [
            Case([1,2,3,0,0,0], 3, [2,5,6], 3, [1,2,2,3,5,6]),
            Case([1], 1, [0], 0, [1]),
            Case([0], 0, [1], 1, [1])
        ]
        
        s = Solution()
        for case in cases:
            s.merge(case.nums1, case.m, case.nums2, case.n)
            self.assertEqual(case.nums1, case.ans)
       

unittest.main()
