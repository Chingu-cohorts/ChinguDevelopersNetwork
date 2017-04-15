defmodule ChinguCentral.Forums.CommentResolver do
  alias ChinguCentral.Repo
  alias ChinguCentral.Forums.Comment
  
  def all(_args, _info) do
    {:ok, Repo.all(Comment)}
  end
end