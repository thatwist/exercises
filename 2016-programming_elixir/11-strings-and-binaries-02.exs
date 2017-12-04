defmodule Anagram do
  def anagram?(word1, word2) do
    Enum.sort(word1) == Enum.sort(word2)
  end
end

IO.inspect(Anagram.anagram?('spar', 'rasp'))
IO.inspect(Anagram.anagram?('spar', 'reps'))
