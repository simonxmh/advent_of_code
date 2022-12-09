#!/usr/bin/env ruby

require 'set'

# load input files
input_file = File.expand_path("../input.txt", __FILE__)
input = File.read(input_file)

arr = []
input.split("\n").each do |line|
    arr.append(line.split("").map(&:to_i))
end

max_score = 0

def calc_score(arr,row,col)
  left = 0
  right = 0
  up =0
  down = 0

  # to up
  (row-1).downto(0).each do |k|
    up += 1
    break if arr[k][col] >= arr[row][col]
  end

  # to down
  (row+1..arr.first.length-1).each do |k|
    down += 1
    break if arr[k][col] >= arr[row][col]
  end

  # to left
  (col-1).downto(0).each do |k|
    left += 1
    break if arr[row][k] >= arr[row][col]
  end

  # to right
  (col+1..arr.length-1).each do |k|
    right += 1
    break if arr[row][k] >= arr[row][col]
  end

  p "#{row}-#{col} #{left}, #{right}, #{up}, #{down}"
  left * right * up * down
end

# from left to right
arr.each_with_index do |line, row|
    line.each_with_index do |elem, col|
        max_score = [calc_score(arr,row,col),max_score].max
    end    
end

p max_score

