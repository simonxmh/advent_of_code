#!/usr/bin/env ruby

# load input files
input_file = File.expand_path("../day_2_input.txt", __FILE__)
input = File.read(input_file)

# start of solution
strat = input.split("\n").map(&:to_str)

win_score = {
    'X' => 0, # lose
    'Y' => 3, # draw
    'Z' => 6  # win
}

shape_score = {
    'A X' => 3, # you choose scissors
    'A Y' => 1, # rock
    'A Z' => 2,
    'B X' => 1,
    'B Y' => 2,
    'B Z' => 3,
    'C X' => 2, # paper
    'C Y' => 3, # scissr
    'C Z' => 1 # rock
}

total = 0

strat.each_with_index do |x,i|
    total += shape_score[x] + win_score[x.split(" ")[1]]
end

puts total
