defmodule Finder do
  def gcd(x, 0), do: x
  def gcd(x, y), do: gcd(y, rem(x, y))
end

IO.puts(Finder.gcd(50, 20))
