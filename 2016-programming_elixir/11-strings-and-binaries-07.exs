defmodule Shop do
  def orders(stream) do
    Enum.map(stream, fn(line) ->
      [id, ship_to, net_amount] = String.split line, ","
      {id, _} = Integer.parse(id)
      ship_to = String.to_atom(String.slice(ship_to, 1..2))
      {net_amount, _} = Float.parse(String.strip net_amount)
      [id: id, ship_to: ship_to, net_amount: net_amount]
    end)
  end

  defp sales_tax(orders) do
    tax_rates = [ NC: 0.075, TX: 0.08]
    Enum.map(orders, fn (order) ->
      net = order[:net_amount]
      tax = @tax_rates[order[:ship_to]] || 0
      order ++ [total_amount: net + net * tax]
    end)
  end
end

{:ok, file} = File.open("orders.csv", [:read])
stream = IO.stream(file, :line)
IO.inspect(Shop.orders(stream))
File.close(file)
