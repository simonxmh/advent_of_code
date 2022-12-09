#!/usr/bin/env ruby

require 'set'

# load input files
input_file = File.expand_path("../input.txt", __FILE__)
input = File.read(input_file)

arr = []
input.split("\n").each do |line|
    arr.append(line.split("").map(&:to_i))
end

included = {}

# from left to right
arr.each_with_index do |line, y|
    max = -1
    line.each_with_index do |elem, x|
        if elem > max
            max = elem
            if !included[ [x,y] ]
                included[ [x,y] ] = true
            end
        end
    end    
end


# from right to left
arr.each_with_index do |line, y|
    max = -1
    line.reverse.each_with_index do |elem, x|
        if elem > max
            max = elem
            if !included[ [line.length() - x - 1, y] ]
                included[ [line.length() - x - 1, y] ] = true
            end
        end
    end    
end


# from top to bottom

arr.transpose.each_with_index do |line, x|
    max = -1
    line.each_with_index do |elem, y|
        if elem > max
            max = elem
            if !included[ [x,y] ]
                included[ [x,y] ] = true
            end
        end
    end    
end

# from bottom to top

arr.transpose.each_with_index do |line, x|
    max = -1
    line.reverse.each_with_index do |elem, y|
        if elem > max
            max = elem
            if !included[ [x, line.length() - y - 1] ]
                included[ [x, line.length() - y - 1] ] = true
            end
        end
    end    
end

p included.sort
p included.length