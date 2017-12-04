defmodule MyString do
  def capitalize_sentence(sentences) do
    sentences 
    |> String.split(". ")
    |> Enum.map(&(String.capitalize &1))
    |> Enum.join(". ")
  end
end

IO.inspect MyString.capitalize_sentence("oh. a DOG. woof. ")
