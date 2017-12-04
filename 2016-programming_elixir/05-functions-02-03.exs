fizzbuzz = fn
  (0, 0, _) -> "FizzBuzz"
  (0, _, _) -> "Fizz"
  (_, 0, _) -> "Buzz"
  (_, _, x) -> x
end

fb = fn (n) -> fizzbuzz.(rem(n, 3), rem(n, 5), n) end

IO.inspect(fb.(10))
IO.inspect(fb.(11))
IO.inspect(fb.(12))
IO.inspect(fb.(13))
IO.inspect(fb.(14))
IO.inspect(fb.(15))
IO.inspect(fb.(16))
