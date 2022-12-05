#!/usr/bin/env ruby

# load input files
input_file = File.expand_path("../day_5_input.txt", __FILE__)
input = File.read(input_file)

# start of solution, do not want to parse this ...
test_stack = [
     ["Z","N"],
     ["M","C","D"],
     ["P"],
]

stack = [
     ["Q","W","P","S","Z","R","H","D"],
     ["V","B","R","W","Q","H","F"],
     ["C","V","S","H"],
     ["H","F","G"],
     ["P","G","J","B","Z"],
     ["Q","T","J","H","W","F","L"],
     ["Z","T","W","D","L","V","J","N"],
     ["D","T","Z","C","J","G","H","F"],
     ["W","P","V","M","B","H"]
]


procedures = input.split("\n\n")[1].split("\n")

procedures.each do |line|
     split_line = line.split(" ")

     amt = split_line[1].to_i
     from = split_line[3].to_i-1
     to = split_line[-1].to_i-1

     tmp = []

     amt.times do
          tmp.push(stack[from].pop())
          stack[to].push(tmp.pop())
     end

end

res = stack.map(&:last).join()

p res
