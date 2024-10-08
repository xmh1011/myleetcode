#include <stdio.h>
#include <stdlib.h>

// 定义链表节点结构体
struct ListNode {
    int val;
    struct ListNode *next;
};

// 函数声明
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    struct ListNode *head = NULL, *tail = NULL;
    int carry = 0;
    while (l1 || l2) {
        int n1 = l1 ? l1->val : 0;
        int n2 = l2 ? l2->val : 0;
        int sum = n1 + n2 + carry;

        // 创建新节点
        if (!head) {
            head = tail = malloc(sizeof(struct ListNode));
            if (!tail) return NULL;  // 检查内存分配是否成功
            tail->val = sum % 10;
            tail->next = NULL;
        } else {
            tail->next = malloc(sizeof(struct ListNode));
            if (!tail->next) return NULL;  // 检查内存分配是否成功
            tail->next->val = sum % 10;
            tail = tail->next;
            tail->next = NULL;
        }

        carry = sum / 10;

        if (l1) l1 = l1->next;
        if (l2) l2 = l2->next;
    }

    // 如果最后还有进位，则创建新节点
    if (carry > 0) {
        tail->next = malloc(sizeof(struct ListNode));
        if (!tail->next) return NULL;  // 检查内存分配是否成功
        tail->next->val = carry;
        tail->next->next = NULL;
    }

    return head;
}

// 创建链表
struct ListNode* createList(int* arr, int size) {
    struct ListNode* head = NULL;
    struct ListNode* tail = NULL;
    for (int i = 0; i < size; ++i) {
        struct ListNode* newNode = malloc(sizeof(struct ListNode));
        if (!newNode) return NULL;  // 检查内存分配是否成功
        newNode->val = arr[i];
        newNode->next = NULL;
        if (!head) {
            head = tail = newNode;
        } else {
            tail->next = newNode;
            tail = newNode;
        }
    }
    return head;
}

// 打印链表
void printList(struct ListNode* node) {
    while (node) {
        printf("%d", node->val);
        node = node->next;
        if (node) printf(" -> ");
    }
    printf("\n");
}

// 释放链表内存
void freeList(struct ListNode* node) {
    while (node) {
        struct ListNode* temp = node;
        node = node->next;
        free(temp);
    }
}

int main() {
    // 创建两个测试链表
    int arr1[] = {2, 4, 3};  // 第一个数 342
    int arr2[] = {5, 6, 4};  // 第二个数 465

    struct ListNode* l1 = createList(arr1, 3);
    struct ListNode* l2 = createList(arr2, 3);

    printf("Input List 1: ");
    printList(l1);

    printf("Input List 2: ");
    printList(l2);

    // 相加两个链表
    struct ListNode* result = addTwoNumbers(l1, l2);

    printf("Result List: ");
    printList(result);

    // 释放内存
    freeList(l1);
    freeList(l2);
    freeList(result);

    return 0;
}
