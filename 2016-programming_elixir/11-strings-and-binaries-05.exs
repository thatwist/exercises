defmodule Strings do
  def center(strings) do
    max = 
      strings
      |> Enum.map(fn(x) -> String.length(x) end)
      |> Enum.max
    for str <- strings do
      IO.puts String.rjust(str, max - (div(max - String.length(str), 2)))
    end
  end
end


Strings.center(["cat", "zebra", "elephant"])
