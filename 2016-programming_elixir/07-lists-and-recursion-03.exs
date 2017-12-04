defmodule MyList do
  def caesar(chars, n) do
    Enum.map(chars, &char(&1 + n))
  end

  defp char(x) when x > 122, do: 97 + x - 123
  defp char(x), do: x
end

IO.puts(MyList.caesar('ryvkve', 13))
