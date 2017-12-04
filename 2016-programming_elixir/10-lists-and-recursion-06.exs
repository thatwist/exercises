defmodule MyEnum do
  def flatten(collection), do: _flatten([], collection)
  def _flatten(acc, []), do: acc
  def _flatten(acc, [head|tail]), do: _flatten(acc, head) ++ _flatten(acc, tail) 
  def _flatten(acc, x), do: acc ++ [x]
end

IO.inspect(MyEnum.flatten([1, 2, [3, 4, [5, 6], [[[7]]], 8]]))
