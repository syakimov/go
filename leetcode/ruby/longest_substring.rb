# https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

# @param {String} s
# @return {Integer}
def length_of_longest_substring(s)
  return 0 if s.size.zero?

  map = {}
  max = 0
  j = 0

  s.each_char.each_with_index do |char, i|
    j = [j, map[char] + 1].max if map[char]
    map[char] = i
    max = [max, i - j + 1].max
  end

  max
end

# @param {String} s
# @return {Integer}
def length_of_longest_substring_stolen(s)
  map = Hash.new(-1)
  start = 0
  ans = 0
  s.chars.each_with_index do |c, index|
    if map[c] >= start
      start = map[c] + 1
    else
      ans = [ans, index - start + 1].max
    end
    map[c] = index
  end
  ans
end

p "Expected 3 from abcabcbb" if length_of_longest_substring('abcabcbb') != 3
p "Expected 1 from bbbbb" if length_of_longest_substring('bbbbb') != 1
p "Expected 3 from pwwkew" if length_of_longest_substring('pwwkew') != 3
p "Expected 5 from tmmzuxt" if length_of_longest_substring('tmmzuxt') != 5
p "Expected 3 from dvdf" if length_of_longest_substring('dvdf') != 3
p "Expected 2 from aab" if length_of_longest_substring('aab') != 2
p "Expected 5 from nfpdmpi" if length_of_longest_substring('nfpdmpi') != 5
