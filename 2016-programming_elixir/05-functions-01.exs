list_concat = fn (x, y) -> x ++ y end
sum = &(&1 + &2 + &3)
pair_tuple_to_list = fn ({x, y}) -> [x, y] end

IO.inspect(list_concat.([:a, :b], [:c, :d]))
IO.inspect(sum.(1, 2, 3))
IO.inspect(pair_tuple_to_list.( {1234, 5678} ))
