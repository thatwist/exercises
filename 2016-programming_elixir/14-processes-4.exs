defmodule Finder do
  def count(file, word) do
    text = File.read!(file)
    Regex.scan(~r/(^|\W)#{word}($|\W)/, text)
    |> Enum.count

  end

  def run(file) do
    receive do
      word -> 2
    end
  end
end

defmodule Scheduler do
  def start(word) do
    File.ls!
    |> Enum.map(fn (file) ->
      spawn(Finder, :run, [file])
    end)
    |> Enum.map(fn (pid) -> send(pid, word) end)
  end
end

IO.inspect Scheduler.start("do")
