#!/usr/bin/env ruby

# load input files
input_file = File.expand_path("../day_4_input.txt", __FILE__)
input = File.read(input_file)

# start of solution
elves = input.split("\n").map(&:to_str)

total = 0
elves.each do |assignments|
     pairs = assignments.split(",")
     elf1, elf2 = pairs[0].split("-").map(&:to_i), pairs[1].split("-").map(&:to_i)

     if elf1[0] < elf2[0] and elf1[1] < elf2[0]
          next
     elsif elf2[0] < elf1[0] and elf2[1] < elf1[0]
          next
     end
     total+=1
     # if elf1[0] <= elf2[0] and elf1[1] >= elf2[1] # elf1 encapsulates elf2
     #      total+=1
     # elsif elf2[0] <= elf1[0] and elf2[1] >= elf1[1] #elf2 encapsulates elf1
     #      total+=1
     # elsif elf1[0] <= elf2[0] and elf1[1] >= elf2[0]  # elf1 end after elf2 start
     #      total+=1
     # elsif elf2[0] <= elf1[0] and elf2[1] >= elf1[0] # elf2 end after elf1 start
     #      total==1
     # end
end

puts total