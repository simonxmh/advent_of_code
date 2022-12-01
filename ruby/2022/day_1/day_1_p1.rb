#!/usr/bin/env ruby


input_file = File.expand_path("../day_1_input.txt", __FILE__)
input = File.read(input_file)

local_calories = 0
max_carried_calories = 0

calories = input.split("\n").map(&:to_i)


calories.each_with_index do |x,i|
    if x == 0 
        local_calories = 0
        next
    end

    local_calories += x

    max_carried_calories = [max_carried_calories,local_calories].max
end

puts max_carried_calories