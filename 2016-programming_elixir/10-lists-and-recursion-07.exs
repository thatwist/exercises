defmodule MyList do
  def span(to, to),   do: [to]
  def span(from, to), do: [from | span(from+1, to) ]

  def prime(x) do
    for n <- MyList.span(2, x), rem(n, 2) == 0, do: n
  end
end

IO.inspect(MyList.prime(40))
