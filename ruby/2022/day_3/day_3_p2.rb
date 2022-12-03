#!/usr/bin/env ruby

# load input files
input_file = File.expand_path("../day_3_input.txt", __FILE__)
input = File.read(input_file)

# start of solution
sacks = input.split("\n").map(&:to_str)

total = 0
sacks.each_slice(3) do |compartment| 

    common = compartment[0].split('') & compartment[1].split('') & compartment[2].split('')

    common.each do |c|
       if c == c.upcase
            total += c.ord - 'A'.ord + 27
       else
            total += c.ord - 'a'.ord + 1
       end
    end
end

puts total
