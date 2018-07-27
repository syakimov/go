# m, n, count integer
# returns [][] integer
def generate_mines_board(r, c, count)
  board = []
  r.times do
    board << Array.new(c, 0)
  end

  mines = []
  while count > 0
    random_row = rand(r)
    random_col = rand(c)

    if board[random_row][random_col].zero?
      board[random_row][random_col] = -1
      mines << [random_row, random_col]
      count -= 1
    end
  end

  p mines
  mines.each do |mine|
    row, col = mine
    increment_count_around_mines(board, row, col)
  end

  board
end

def increment_count_around_mines(board, row, col)
  [
    [row-1, col-1],
    [row-1, col],
    [row-1, col+1],
    [row, col-1],
    [row, col+1],
    [row+1, col-1],
    [row+1, col],
    [row+1, col+1]
  ].each do |coordinates|
    crow, ccol = coordinates

    if board[crow] && board[crow][ccol] && board[crow][ccol] != -1
      board[crow][ccol] += 1
    end
  end
end

b =  generate_mines_board(5, 4, 4)
b.each do |row|
  p row
end
