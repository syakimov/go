# frozen_string_literal: true

class ListNode
  attr_accessor :val, :next
  def initialize(val, n = nil)
    @val = val
    @next = n
  end
end

# @param {ListNode} l1
# @param {ListNode} l2
# @return {ListNode}
def add_two_numbers(l1, l2)
  carry = 0
  dummy_head = n = ListNode.new(0)

  while !l1.nil? || !l2.nil? || carry != 0
    v1 = l1 ? l1.val : 0
    v2 = l2 ? l2.val : 0

    l1 = l1&.next
    l2 = l2&.next

    v = v1 + v2 + carry

    if v >= 10
      carry = 1
      v = v % 10
    else
      carry = 0
    end

    n.next = ListNode.new(v)
    n = n.next
  end

  dummy_head.next
end

# Tests
x = add_two_numbers(ListNode.new(2, ListNode.new(4, ListNode.new(3, nil))),
                    ListNode.new(5, ListNode.new(6, ListNode.new(4, nil))))

p x.val           # 7
p x.next.val      # 0
p x.next.next.val # 8
