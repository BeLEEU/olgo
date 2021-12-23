
//23找到链表第一个重合点
//定义计数器
//遍历两个链表直到为空，统计链表节点个数
//算出差值，长点的链表先走差值的步数
//一起遍历，直到找到相等的节点
class ListNode {
  val: number
  next: ListNode | null
  constructor(val: number, next: ListNode | null) {
    this.val = val == undefined ? 0 : val;
    this.next = next == null ? null : next;
  }
}

function getInsertionNode(headA: ListNode | null, headB: ListNode | null): ListNode | null {
  
  let countA: number = 0;
  let countB: number = 0;
  let node1: ListNode = headA;
  let node2: ListNode = headB;

  while(node1 != null) {
    countA += 1;
    node1 = node1.next;
  }

  while(node2 != null) {
    countB += 1;
    node2 = node2.next;
  }

  node1 = headA;
  node2 = headB;
  let count = countA - countB;
  if(countA > countB) {
    while(count > 0) {
      node1 = node1.next;
      count -= 1;
    }
  } else {
    while(count < 0) {
      node2 = node2.next;
      count += 1;
    }
  }

  while(node2 != null) {
    if(node1 === node2) {
      return node2;
    }
    node1 = node1.next;
    node2 = node2.next;
  }
  return null;
}