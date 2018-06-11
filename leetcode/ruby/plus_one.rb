require 'pry'

# @param {Integer[]} digits
# @return {Integer[]}
def plus_one(digits)
  i = digits.size - 1
  bonus = 1

  while i >= 0 && bonus.positive?
    bonus = (digits[i] + 1) / 10
    digits[i] = (digits[i] + 1) % 10

    i -= 1
  end

  digits.unshift(1) if digits.first == 0

  digits
end

p plus_one [1, 2, 3]
p plus_one [9, 9, 9]
