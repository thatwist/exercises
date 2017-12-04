defmodule Calc do
  def sum([]), do: 0
  def sum([head|tail]), do: head + sum(tail)
end

IO.inspect(Calc.sum([5, 10, 21]))
