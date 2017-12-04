defmodule Ticker do

  @tick 2000 # 2 seconds
  @name :ticker

  def start() do
    pid = spawn(__MODULE__, :receiver, [[]])
    :global.register_name(@name, pid)
  end

  def register(client_pid) do
    send(:global.whereis_name(@name), {:register, client_pid})
  end

  def receiver(client_pids) do
    receive do
      {:register, client_pid} ->
        IO.puts "Client registered #{inspect client_pid}"
        receiver([client_pid|client_pids])
    after
      @tick ->
        IO.puts("Tick in server")
        Enum.each(client_pids, fn client_pid ->
          send(client_pid, {:tick})
        end)
        receiver(client_pids)
    end
  end
end

defmodule Client do
  def start() do
    pid = spawn(__MODULE__, :receiver, [])
    Ticker.register(pid)
  end

  def receiver do
    receive do
      {:tick} -> IO.puts("Tock in client")
    end
    receiver
  end
end
