defmodule My do
  defmacro times_n(n) do
    quote do
      def unquote(:"times_#{n}")(x) do
        x * unquote(n)
      end
    end
  end
end
