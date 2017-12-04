defmodule MyList do
  def max([head|tail]), do: _max(tail, head)

  defp _max([], m), do: m
  defp _max([head|tail], m) when m >= head, do: _max(tail, m)
  defp _max([head|tail], m) when m < head, do: _max(tail, head)
end

IO.inspect(MyList.max([10, 2, 3, 4, 17, 100, 103, 20, 24]))
