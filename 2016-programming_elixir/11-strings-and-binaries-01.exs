defmodule MyChars do
  def all_printable?(chars) do
    Enum.all?(chars, fn (c) -> c in ?\s..?~ end)
  end
end

IO.inspect(MyChars.all_printable?('elixir'))
IO.inspect(MyChars.all_printable?('ł'))
