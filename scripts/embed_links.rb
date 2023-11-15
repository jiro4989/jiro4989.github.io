#!/usr/bin/ruby

post_links = STDIN.map{|l| l.strip}

is_posts = false
contents = []
file = "index.md"
IO.readlines(file, chomp: true).each{|l|
  if l.include?("START_POSTS")
    is_posts = true
    contents.append(l)
    for link in post_links do
      contents.append(link)
    end
  elsif l.include?("END_POSTS")
    is_posts = false
  end

  if not is_posts
    contents.append(l)
  end
}

f = File.open(file, "w")

for c in contents do
  f.write(c)
  f.write("\n")
end

f.close()
