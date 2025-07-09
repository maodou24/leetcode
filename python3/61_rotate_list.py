# Definition for singly-linked list.
from typing import Optional
import unittest

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        n = 0
        

        node = head
        while node is not None:
            n += 1
            node = node.next
        
        k = k % n

        head, last = head, head
        count = 0
        while count == k:
