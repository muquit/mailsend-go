#!/usr/bin/env ruby

# muquit@muquit.com Jan-20-2019 
# force tag 

require "highline/import"

def check_input(input)
  if input.length == 0
    puts "Nothing specified. Exiting.."
    exit 1
  end
end

begin
  o = `git tag|sort`.chomp
  tags = o.split("\n")
  tags.each do |tag|
    puts "#{tag}"
  end

  last_tag = `git describe --abbrev=0 --tags`
  puts ""
  puts "Last Tag: #{last_tag}"
  puts ""

  tag = ask "Enter Tag: "
  check_input(tag)
  puts "Tag: #{tag}"

  comment = "Release #{tag}"
  puts "Comment: #{comment}"

  cmd = "git tag -f -a #{tag} -m \"#{comment}\""
  puts "> #{cmd}"

  ans = ask "Press Enter to continue: "
  if ans.length > 0
    puts "Exit..."
    exit 1
  end

  system(cmd)
  cmd = "git push -f origin #{tag}"
  puts "> #{cmd}"
  system(cmd)
  

rescue => e
  puts "ERROR: #{e}"
end
