# Definition for singly-linked list.
from typing import List


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __eq__(self, other): 
        if not isinstance(other, ListNode):
            return NotImplemented

        node1, node2 = self, other
        while node1 is not None and node2 is not None:
            if node1.val != node2.val:
                return False
            node1 = node1.next
            node2 = node2.next
        return node1 is None and node2 is None


def to_listNode(list: List[int]) -> ListNode:
    if len(list) == 0:
        return None
    
    head = ListNode()
    node = head
    for v in list:
        node.next = ListNode(v)
        node = node.next
    
    return head.next
