#!/usr/bin/env ruby

# load input files
input_file = File.expand_path("../day_2_input.txt", __FILE__)
input = File.read(input_file)

# start of solution
strat = input.split("\n").map(&:to_str)

shape_score = {
    'X' =>1,
    'Y' =>2,
    'Z' =>3
}

winning_relations = {
    'A X' => 3,
    'A Y' => 6,
    'A Z' => 0,
    'B X' => 0,
    'B Y' => 3,
    'B Z' => 6,
    'C X' => 6,
    'C Y' => 0,
    'C Z' => 3
}

total = 0

strat.each_with_index do |x,i|
    shape = shape_score[x.split(" ")[1]]
    total += shape + winning_relations[x]
end

puts total
