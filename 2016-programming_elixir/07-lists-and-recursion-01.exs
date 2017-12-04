defmodule MyList do
  def mapsum([], f), do: 0
  def mapsum([head|tail], f), do: f.(head) + mapsum(tail, f)
end

IO.inspect(MyList.mapsum [1, 2, 3], &(&1 * &1))
