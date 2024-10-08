#include <iostream>
using namespace std;

// 定义链表节点结构体
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *head = nullptr, *tail = nullptr;
        int carry = 0;
        while (l1 || l2) {
            int n1 = l1 ? l1->val : 0;
            int n2 = l2 ? l2->val : 0;
            int sum = n1 + n2 + carry;

            // 创建新的节点存储当前位的值
            if (!head) {
                head = tail = new ListNode(sum % 10);
            } else {
                tail->next = new ListNode(sum % 10);
                tail = tail->next;
            }

            carry = sum / 10;
            if (l1) l1 = l1->next;
            if (l2) l2 = l2->next;
        }

        // 如果最高位有进位，则需要新建节点
        if (carry > 0) {
            tail->next = new ListNode(carry);
        }
        return head;
    }
};

// 打印链表的内容
void printList(ListNode* node) {
    while (node) {
        cout << node->val;
        node = node->next;
        if (node) cout << " -> ";
    }
    cout << endl;
}

// 创建链表
ListNode* createList(const initializer_list<int>& values) {
    ListNode* head = nullptr;
    ListNode* tail = nullptr;
    for (int value : values) {
        if (!head) {
            head = tail = new ListNode(value);
        } else {
            tail->next = new ListNode(value);
            tail = tail->next;
        }
    }
    return head;
}

int main() {
    Solution solution;

    // 创建两个链表: l1 = 2 -> 4 -> 3, l2 = 5 -> 6 -> 4
    ListNode* l1 = createList({2, 4, 3});
    ListNode* l2 = createList({5, 6, 4});

    cout << "Input List 1: ";
    printList(l1);

    cout << "Input List 2: ";
    printList(l2);

    // 相加链表
    ListNode* result = solution.addTwoNumbers(l1, l2);

    cout << "Result List: ";
    printList(result);

    // 释放链表内存
    while (l1) {
        ListNode* temp = l1;
        l1 = l1->next;
        delete temp;
    }

    while (l2) {
        ListNode* temp = l2;
        l2 = l2->next;
        delete temp;
    }

    while (result) {
        ListNode* temp = result;
        result = result->next;
        delete temp;
    }

    return 0;
}
