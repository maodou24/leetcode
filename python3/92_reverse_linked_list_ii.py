# Definition for singly-linked list.
from typing import Optional
import unittest

from common import *
        
class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        dummy_node = ListNode(-1)
        dummy_node.next = head
        pre = dummy_node
        for _ in range(left - 1):
            pre = pre.next

        cur = pre.next
        for _ in range(right - left):
            next = cur.next
            cur.next = next.next
            next.next = pre.next
            pre.next = next

        return dummy_node.next


class SolutionTestCase(unittest.TestCase):

    def test_reverseBetween(self):
        class Case:
            def __init__(self, head: Optional[ListNode], left, right, ans: Optional[ListNode]):
                self.head = head
                self.left = left
                self.right = right
                self.ans = ans
        
        cases = [
            Case([1, 2, 3, 4, 5], 2, 4, [1,4,3,2,5]),
            Case([3, 5], 1, 2, [5, 3]),
        ]
        
        s = Solution()
        for case in cases:
            self.assertEqual(s.reverseBetween(to_listNode(case.head), case.left, case.right), to_listNode(case.ans))


unittest.main()
