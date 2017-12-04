defmodule My do
  defmacro unless(clause, expression) do
    do_exp = case Keyword.fetch(expression, :do) do
      {:ok, block} -> block
      _ -> raise ArgumentError, message: "invalid keys for unless, only 'do' is permitted"
    end
    quote do
      if(!unquote(clause), do: unquote(do_exp))
    end
  end
end
