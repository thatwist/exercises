defmodule Mod do
  def sum(n, acc \\ 0)

  def sum(0, acc), do: acc
  def sum(n, acc), do: sum(n - 1, acc + n)
end
