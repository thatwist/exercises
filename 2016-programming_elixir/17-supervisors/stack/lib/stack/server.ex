defmodule Stack.Server do
  use GenServer

  # Client API

  def start_link(init) do
    GenServer.start_link(__MODULE__, init, [name: __MODULE__])
  end

  def pop do
    GenServer.call(__MODULE__, {:pop})
  end

  # GenServer implementation

  def handle_call({:pop}, _caller, stack_state) do
    [head | tail] = stack_state
    {:reply, head, tail}
  end
end
