defmodule Parallel do
  def pmap(collection, f) do
    parent_pid = self
    collection
    |> Enum.map(fn (elem) ->
      spawn_link(fn -> send(parent_pid, {self, f.(elem)}) end)
    end)
    |> Enum.map(fn (pid) ->
      receive do
        {^pid, value} -> value
      end
    end)
  end
end

IO.inspect Parallel.pmap(1..1000, fn (x) -> x + 1 end)
