defmodule Finder do
  def count(file, word) do
    text = File.read!(file)
    Regex.scan(~r/(^|\W)#{word}($|\W)/, text)
    |> Enum.count
  end

  def run(parent_pid, file) do
    receive do
      word -> send(parent_pid, {file, count(file, word)})
    end
  end
end

defmodule Scheduler do
  def start(word) do
    File.ls!
    |> Enum.map(fn (file) ->
      spawn(Finder, :run, [self, file])
    end)
    |> Enum.map(fn (pid) -> send(pid, word) end)
    |> Enum.map(fn (_pid) ->
      receive do
        {file, word} -> {file, word}
      end
    end)
  end
end

IO.inspect Scheduler.start("receive")
