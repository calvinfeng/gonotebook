target = 100_000_000
num_of_threads = 5
threads = []

num_of_threads.times do
  threads << Thread.new do
    sum = 0
    (target/num_of_threads).times do
      sum += 1
    end
    Thread.current[:output] = sum
  end
end

start_time = Time.new
total = 0
threads.each do |thr|
  thr.join
  total += thr[:output]
end
end_time = Time.new

puts "The total sum is #{total}"
puts end_time - start_time
