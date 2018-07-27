# Sliven -> Burgas
# Plovdiv -> Stara Zagora
# Stara Zagora -> Sliven
# Burgas -> Varna
# Sofia -> Plovdiv


# Cycle -> no
# Streight line
# Start
# Next
# Finish?


# Sofia -> Plovdiv

# Plovdiv -> Stara_Zagora
# Stara_Zagora -> Sliven
# Sliven -> Burgas
# Burgas -> Varna


# 1. Find Single city at first place -> 0(n)
#
# 2.1. Remove line from list
# 2.2. Find next start point
# 3. If List count == 0


# Sofia -> Plovdiv
# Print
# hash.delete(key) // key = startpoint

require 'set'
require 'pry'

def sort_tickets(tickets)
  return nil if tickets.nil?

  hash = {}
  tickets.each do |ticket|
    key, value =  ticket.split(' -> ')
    hash[key] = value
  end

  start = (hash.keys - hash.values).first # set - set -> n

  while hash.keys.size != 0
    destination = hash[start]
    p "#{start} -> #{destination}"
    hash.delete(start)
    start = destination
  end
end

sort_tickets([
  'Sliven -> Burgas',
  'Plovdiv -> Stara Zagora',
  'Stara Zagora -> Sliven',
  'Burgas -> Varna',
  'Sofia -> Plovdiv' ])
