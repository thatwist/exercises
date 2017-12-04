defmodule FizzBuzz do
  def fizzbuzz(n) do
    case {rem(n, 3), rem(n, 5)} do
      {0, 0} -> "FizzBuzz"
      {0, _} -> "Fizz"
      {_, 0} -> "Buzz"
      {_, _} -> n
    end
  end
end

Enum.each(1..100, fn(n) ->
  IO.inspect(FizzBuzz.fizzbuzz(n)) end)
