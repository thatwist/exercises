defmodule Incrementer do
  def start(next_pid) do
    receive do
      i -> send(next_pid, i + 1)
    end
  end

  def run(n) do
    last_pid = 1..n |> Enum.reduce(self, fn (_i, acc) ->
      pid = spawn(__MODULE__, :start, [acc])
      pid
    end)

    send(last_pid, 0)

    receive do
      i -> IO.puts("Got #{i} as a result")
    after
      3000 ->
        IO.puts :stderr, "No message in 3 seconds"
    end
  end
end

IO.inspect :timer.tc(Incrementer, :run, [1_000_000])
