#!/usr/bin/env ruby

# load input files
input_file = File.expand_path("../day_7_input.txt", __FILE__)
input = File.read(input_file)

directory_sizes = Hash.new(0)
cwd = []

input.split("\n").each do |cmd|
    if cmd == "$ cd /"
        cwd = []
    elsif cmd == "$ cd .."
        cwd.pop()
    elsif cmd.split(" ")[..1] == ["$","cd"]
        cwd.append(cmd.split(" ")[2])
    elsif cmd == "$ ls"
        #noop
    elsif cmd.split(" ")[0] == "dir"
        #noop
    else
        p cwd
        cwd.each do |prefix|
            directory_sizes[prefix] += cmd.split(" ")[0].to_i
        end
    end

    p directory_sizes
end 
p directory_sizes
p directory_sizes.select{|k,v| v<=100000}.values.sum