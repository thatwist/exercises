prefix = fn (x) -> fn (y) -> "#{x} #{y}" end end

mrs = prefix.("Mrs.")
elixir = prefix.("Elixir")

IO.inspect(mrs.("Smith"))
IO.inspect(elixir.("Rocks!"))
