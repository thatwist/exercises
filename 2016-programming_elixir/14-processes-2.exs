defmodule Tokens do
  def start do
    pid1 = spawn(Tokens, :reply, [self])
    pid2 = spawn(Tokens, :reply, [self])

    send(pid2, "betty")
    send(pid1, "fred")

    receive_token(2)
  end

  def reply(pid) do
    receive do
      token -> send(pid, token)
    end
  end

  defp receive_token(0), do: nil
  defp receive_token(n) do
    receive do
      token -> IO.puts token
    end
    receive_token(n-1)
  end
end

Tokens.start
