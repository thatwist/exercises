defmodule Ok do
  def ok!({:ok, result}), do: result
  def ok!({err, result}), do: raise(RuntimeError, "Not ok! Got {#{err}, #{result}}")
end

Ok.ok! File.open("orders.csv")
Ok.ok! File.open("test.log")
