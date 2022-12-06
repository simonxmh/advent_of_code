#!/usr/bin/env ruby

# load input files
input_file = File.expand_path("../day_6_input.txt", __FILE__)
input = File.read(input_file)


input.split("").each_with_index do |c,i|
     p input[i-13..i].split("").uniq
     if input[i-13..i].split("").uniq.length == 14
          p i+1
          break
     end
end