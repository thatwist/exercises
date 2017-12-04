defmodule Calculator do
  def call(chars) do
    [o1, op, o2] = String.split String.Chars.to_string(chars)
    operate(op, parse_int(o1), parse_int(o2))
  end

  defp parse_int(str) do
    {int, dec} = Integer.parse(str)
    int
  end

  defp operate("+", o1, o2), do: o1 + o2
  defp operate("-", o1, o2), do: o1 - o2
  defp operate("*", o1, o2), do: o1 * o2
  defp operate("/", o1, o2), do: o1 / o2
end

IO.inspect(Calculator.call('23 + 48'))
IO.inspect(Calculator.call('87 - 2'))
IO.inspect(Calculator.call('24 * 2'))
IO.inspect(Calculator.call('12 / 4'))
