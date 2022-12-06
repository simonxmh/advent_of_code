#!/usr/bin/env ruby

# load input files
input_file = File.expand_path("../day_6_input.txt", __FILE__)
input = File.read(input_file)


input[3..].split("").each_with_index do |c,i|
     p input[i-3..i].split("").uniq
     if input[i-3..i].split("").uniq.length == 4
          p i+1
          break
     end
end