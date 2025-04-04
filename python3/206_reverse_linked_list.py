
from typing import Optional
import unittest
from common import *


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return head
        
        prev, next = head, head.next
        while next is not None:
            temp = next.next
            next.next = prev

            prev = next
            next = temp

        head.next = None
        return prev
        
class SolutionTestCase(unittest.TestCase):

    def test_reverseList(self):
        class Case:
            def __init__(self, head: Optional[ListNode], ans: Optional[ListNode]):
                self.head = head
                self.ans = ans
        
        cases = [
            Case([1, 2, 3, 4, 5], [5,4,3,2,1]),
            Case([1,2], [2,1]),
            Case([], [])
        ]
        
        s = Solution()
        for case in cases:
            self.assertEqual(s.reverseList(to_listNode(case.head)), to_listNode(case.ans))


unittest.main()