# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
from typing import Optional

from python3.common import ListNode


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        carry = int(0)
        header = ListNode
        node = header
        while l1 is not None or l2 is not None:
            sum = carry

            if l1 is not None:
                sum += l1.val
                l1 = l1.next

            if l2 is not None:
                sum += l2.val 
                l2 = l2.next
   
            carry = int(sum / 10)
            node.next = ListNode(val=sum%10)
            node = node.next

        if carry > 0:
            node.next = ListNode(val=carry)

        return header.next