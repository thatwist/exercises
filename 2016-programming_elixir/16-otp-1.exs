defmodule Stack do
  use GenServer

  # Client API

  def start_link(items, options \\ []) do
    GenServer.start_link(__MODULE__, items, options)
  end

  def pop(pid) do
    GenServer.call(pid, {:pop})
  end

  def push(pid, item) do
    GenServer.cast(pid, {:push, item})
  end

  # Server implementation

  def handle_call({:pop}, _caller, items) do
    [head | tail] = items
    {:reply, head, tail}
  end

  def handle_cast({:push, item}, items) do
    {:noreply, [item | items]}
  end
end
