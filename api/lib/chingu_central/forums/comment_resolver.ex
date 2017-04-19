defmodule ChinguCentral.Forums.CommentResolver do
  alias ChinguCentral.Repo
  alias ChinguCentral.Forums.Comment
  
  def all(_args, _info) do
    {:ok, Repo.all(Comment)}
  end
  
  def find(%{id: id}, _info) do
    case Repo.get(Comment, id) do
      nil -> {:error, "Comment with ID #{id} not found"}
      comment -> {:ok, comment}
    end
  end
end