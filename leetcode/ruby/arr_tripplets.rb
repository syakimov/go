require 'pry'

@res = []

def help(arr, target, temp, index)
  if temp.size == 3
    @res.push [].concat temp if temp.inject(0, :+) > target
    return
  end

  while index < arr.size
    temp.push(arr[index])
    help(arr, target, temp, index + 1)
    temp.pop
    index += 1
  end
end

def tripplets(arr, val)
  help(arr, val, [], 0)

  p @res.sort.uniq
end

tripplets([1, 2, 3, 4, 5], 9)
# [1, 4, 5]
# [2, 3, 5]
# [2, 4, 5]
# [3, 4, 5]
