use std::option::Option;
use std::fmt;

// 定义链表节点结构体
#[derive(Debug)]
struct ListNode {
    val: i32,
    next: Option<Box<ListNode>>,
}

impl ListNode {
    // 创建一个新的节点
    fn new(val: i32) -> Self {
        ListNode { val, next: None }
    }

    // 在链表末尾添加新节点
    fn append(&mut self, val: i32) {
        let mut curr = self;
        while let Some(ref mut next) = curr.next {
            curr = next;
        }
        curr.next = Some(Box::new(ListNode::new(val)));
    }
}

// 实现链表的打印功能
impl fmt::Display for ListNode {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let mut curr = Some(self);
        while let Some(node) = curr {
            write!(f, "{} ", node.val)?;
            curr = node.next.as_deref();
            if curr.is_some() {
                write!(f, "-> ")?;
            }
        }
        Ok(())
    }
}

struct Solution;

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut l1 = l1;
        let mut l2 = l2;
        let mut head = None;
        let mut tail = &mut head;
        let mut carry = 0;

        // 遍历两个链表
        while l1.is_some() || l2.is_some() {
            let n1 = l1.as_ref().map_or(0, |node| node.val);
            let n2 = l2.as_ref().map_or(0, |node| node.val);
            let sum = n1 + n2 + carry;

            // 计算进位和当前位的值
            carry = sum / 10;

            // 创建新的节点存储当前位的值
            let new_node = Box::new(ListNode::new(sum % 10));
            match tail {
                Some(node) => node.next = Some(new_node),
                None => *tail = Some(new_node),
            }
            tail = &mut tail.as_mut().unwrap().next;

            // 移动指针
            l1 = l1.and_then(|node| node.next);
            l2 = l2.and_then(|node| node.next);
        }

        // 如果最高位有进位，则需要新建节点
        if carry > 0 {
            let new_node = Box::new(ListNode::new(carry));
            match tail {
                Some(node) => node.next = Some(new_node),
                None => *tail = Some(new_node),
            }
        }

        head
    }
}

// 创建链表
fn create_list(values: &[i32]) -> Option<Box<ListNode>> {
    let mut head = None;
    let mut tail = &mut head;
    for &value in values {
        let new_node = Box::new(ListNode::new(value));
        match tail {
            Some(node) => node.next = Some(new_node),
            None => *tail = Some(new_node),
        }
        tail = &mut tail.as_mut().unwrap().next;
    }
    head
}

fn main() {
    // 创建两个链表: l1 = 2 -> 4 -> 3, l2 = 5 -> 6 -> 4
    let l1 = create_list(&[2, 4, 3]);
    let l2 = create_list(&[5, 6, 4]);

    // 打印输入链表
    if let Some(ref node) = l1 {
        println!("Input List 1: {}", node);
    }
    if let Some(ref node) = l2 {
        println!("Input List 2: {}", node);
    }

    // 相加链表
    let result = Solution::add_two_numbers(l1, l2);

    // 打印结果链表
    if let Some(ref node) = result {
        println!("Result List: {}", node);
    }
}
