defmodule MyEnum do
  # all?

  def all?([], _f), do: true
  def all?([head | tail], f) do
    if (f.(head)) do
      all?(tail, f)
    else
      false
    end
  end

  # each
  #
  def each([], _f), do: nil
  def each([head|tail], f) do
    f.(head)
    each(tail, f)
  end

  # filter
  #
  def filter([], _f), do: []
  def filter([head|tail], f) do
    if (f.(head)) do
      [head] ++ filter(tail, f)
    else
      filter(tail, f)
    end
  end

  # split
  #
  def split(a, count) when count >= 0, do: _split({[], a}, count)
  def split(a, count) when count  < 0, do: _split({[], a}, _length(a) + count)

  def _split({x, y}, count) when count <= 0, do: {x, y}
  def _split({x, []}, count), do: {x, []}
  def _split({x, [yh|yt]}, count), do: _split({x ++ [yh], yt}, count - 1)

  def _length(x), do: _length(x, 0)
  def _length([], count), do: count
  def _length([head|tail], count), do: _length(tail, count + 1)

  # take
  #
  def take(collection, count) when count >= 0, do: _take([], collection, count)
  def take(collection, count) when count < 0,  do: _take([], _reverse([], collection), -count)
  def _take(acc, _collection, 0), do: acc
  def _take(acc, [], count), do: acc
  def _take(acc, [head|tail], count), do: _take(acc ++ [head], tail, count - 1)
  def _reverse(acc, []), do: acc
  def _reverse(acc, [head|tail]), do: _reverse([head] ++ acc, tail)
end

IO.puts "\n## all?"
IO.inspect(MyEnum.all?([1, 2, 3], &(&1 < 5)))

IO.puts "\n## each"
MyEnum.each([1, 2, 3], &(IO.puts(&1)))

IO.puts "\n## filter"
IO.inspect(MyEnum.filter([7, 20, 4, 3, 17], &(&1 < 10)))

IO.puts "\n## split"
IO.inspect(MyEnum.split([1, 2, 3], 2))
IO.inspect(MyEnum.split([1, 2, 3], 10))
IO.inspect(MyEnum.split([1, 2, 3], 0))
IO.inspect(MyEnum.split([1, 2, 3], -1))
IO.inspect(MyEnum.split([1, 2, 3], -5))

IO.puts "\n## take"
IO.inspect(MyEnum.take([1, 2, 3], 2))
IO.inspect(MyEnum.take([1, 2, 3], 10))
IO.inspect(MyEnum.take([1, 2, 3], 0))
IO.inspect(MyEnum.take([1, 2, 3], -1))
