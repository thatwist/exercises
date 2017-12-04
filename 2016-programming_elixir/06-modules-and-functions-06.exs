defmodule Chop do
  def guess(actual, a..b), do: guessing(actual, make_a_guess(a, b), a..b)

  defp guessing(actual, actual, _a.._b), do:
    IO.puts(actual)
  defp guessing(actual, guess, _a..b) when actual > guess, do:
    guessing(actual, make_a_guess(guess, b), guess..b)
  defp guessing(actual, guess, a.._b) when actual < guess, do:
    guessing(actual, make_a_guess(a, guess), a..guess)

  defp make_a_guess(a, b) do
    x = a + div(b-a, 2)
    IO.puts("Is it #{x}")
    x
  end
end

Chop.guess(273, 1..1000)
