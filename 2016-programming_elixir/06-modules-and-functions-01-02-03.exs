defmodule Times do
  def double(x), do: 2 * x
  def triple(x), do: 3 * x
  def quadruple(x), do: x |> double |> double
end

IO.puts(Times.triple(3))
IO.puts(Times.quadruple(3))
