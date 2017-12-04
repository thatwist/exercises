defmodule Yo do
  import :timer, only: [sleep: 1]

  def ping(parent_pid) do
    send(parent_pid, :yo)
  end

  def start do
    spawn_link(__MODULE__, :ping, [self])
    sleep(500)
    receive do
      x -> IO.puts(x)
    end
  end
end

Yo.start
