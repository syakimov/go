# @param {Integer[]} nums
# @param {Integer} k
# @return {Void} Do not return anything, modify nums in-place instead.
def rotate(nums, k)
  k %= nums.size
  (nums.size - k).times do
    nums.push(nums.shift)
  end
end

# p rotate([1,2,3,4,5,6,7], 3)
# # [5,6,7,1,2,3,4]
# p rotate([1,2], 2)
# # [1, 2]
# p rotate([1,2], 3)
# # [2,1]
# p rotate([1,2,3], 1)
# # [3,1,2]
