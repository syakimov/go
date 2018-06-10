# @param {String} s
# @return {Boolean}
def is_palindrome(s)
  s = s.downcase.gsub(/[^0-9a-z]/, '')
  i = 0
  while i < s.size / 2
    return false if s[i] != s[-i - 1]
    i += 1
  end

  true
end

p is_palindrome('A man, a plan, a canal: Panama')
p is_palindrome('0P')

def is_palindrome_stolen(s)
  str = s.delete('^0-9a-zA-Z').downcase
  str.reverse == str
end

def is_palindrome_stolen_fastest(s)
  s.downcase!
  s.gsub!(/\W/, '')
  s == s.reverse
end
