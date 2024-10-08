class ListNode {
    int val;          // 节点的值
    ListNode next;    // 指向下一个节点的指针

    ListNode(int val) {
        this.val = val;
        this.next = null;
    }

    // 在链表末尾添加一个新节点
    void append(int val) {
        ListNode curr = this;
        while (curr.next != null) {
            curr = curr.next;
        }
        curr.next = new ListNode(val);
    }

    // 重写 toString 方法，打印链表的内容
    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        ListNode curr = this;
        while (curr != null) {
            sb.append(curr.val);
            if (curr.next != null) {
                sb.append(" -> ");
            }
            curr = curr.next;
        }
        return sb.toString();
    }
}

class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode head = null, tail = null;
        int carry = 0;

        // 遍历两个链表
        while (l1 != null || l2 != null) {
            int n1 = (l1 != null) ? l1.val : 0;
            int n2 = (l2 != null) ? l2.val : 0;
            int sum = n1 + n2 + carry;

            // 计算进位和当前位的值
            carry = sum / 10;

            // 创建新的节点存储当前位的值
            ListNode newNode = new ListNode(sum % 10);
            if (head == null) {
                head = tail = newNode;
            } else {
                tail.next = newNode;
                tail = newNode;
            }

            // 移动指针
            if (l1 != null) l1 = l1.next;
            if (l2 != null) l2 = l2.next;
        }

        // 如果最高位有进位，则需要新建节点
        if (carry > 0) {
            tail.next = new ListNode(carry);
        }

        return head;
    }
}

public class Main {
    // 创建链表
    public static ListNode createList(int[] values) {
        if (values.length == 0) return null;
        ListNode head = new ListNode(values[0]);
        for (int i = 1; i < values.length; i++) {
            head.append(values[i]);
        }
        return head;
    }

    public static void main(String[] args) {
        // 创建两个链表: l1 = 2 -> 4 -> 3, l2 = 5 -> 6 -> 4
        int[] values1 = {2, 4, 3};
        int[] values2 = {5, 6, 4};

        ListNode l1 = createList(values1);
        ListNode l2 = createList(values2);

        // 打印输入链表
        System.out.println("Input List 1: " + l1);
        System.out.println("Input List 2: " + l2);

        // 相加链表
        Solution solution = new Solution();
        ListNode result = solution.addTwoNumbers(l1, l2);

        // 打印结果链表
        System.out.println("Result List: " + result);
    }
}
