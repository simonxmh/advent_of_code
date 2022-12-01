#!/usr/bin/env ruby


input_file = File.expand_path("../day_1_input.txt", __FILE__)
input = File.read(input_file)

local_calories = 0
max_carried_calories = [0,0,0]

calories = input.split("\n").map(&:to_i)

aggregated_calories = []

calories.each_with_index do |x,i|
    if x == 0 
        aggregated_calories.push(local_calories)
        local_calories = 0
        next
    end

    local_calories += x
end

puts aggregated_calories.max(3).sum