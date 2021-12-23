// class ListNode {
//   val: number
//   next: ListNode | null
//   constructor(val?: number, next?: ListNode | null) {
//     this.val = val == undefined ? 0 : val;
//     this.next = next == undefined ? null : next;
//   }
// }
/**
 * //206反转链表
 * @param head 
 * 一、迭代法
 * 1.定义dummy节点，并且next指针指向head
 * 2.初始化pre和cur指针，分别指向dummy和head
 * 3.保存cur指针的next到temp指针
 * 4.cur指针的next指向pre
 * 5.pre指针指向cur
 * 6.cur指针指向temp
 * 7.迭代3-6步，直到cur为null
 * 8.此时pre指向反转之后的头结点，返回pre（清除dummy的next，并将dummy设为null）
 */
function reverseListNode(head: ListNode | null): ListNode | null {
  
  if(head == null) {
    return null;
  }

  let dummy: ListNode | null = new ListNode(0, null);
  dummy.next = head;
  let preNode: ListNode | null = dummy;
  let curNode: ListNode | null = dummy.next;

  while(curNode != null) {
    let tempNode = curNode.next;
    curNode.next = preNode;
    preNode = curNode;
    curNode = tempNode;
  }

  dummy.next = null;
  dummy = null;
  return preNode;
}

/**
 * 
 * @param head 
 * @returns 
 * 二、递归法
 * 1.当当前节点是空节点或者next为null时，直接返回该节点
 * 2.当前节点的下下节点，应该指向当前节点，并且当前节点的next应该置为空
 * 3.如果是从头到尾这么做，会丢失连接的指针，因此需要递归地从尾部开始
 * 4.从尾部开始，因为尾部的next指针为null，所以直接返回尾部节点
 * 5.倒数第二个，将尾部的next指针指向自己，并将自己的next置为空
 * 6.如5步骤执行到头部
 * 7.因为需要返回反转后的头指针，所以最终是需要返回尾部节点的，因此将尾部的返回值传给一个临时变量，依次回传
 */
function reverseListNode1(head: ListNode | null): ListNode | null {
  
  if(head == null || head.next == null) {
    return head;
  }

  let resultNode = reverseListNode1(head.next)
  head = head.next.next;
  head.next = null;
  return resultNode;
  
}