defmodule ChinguCentral.Forums.PostResolver do
  alias ChinguCentral.Repo
  alias ChinguCentral.Forums.Post
  
  def all(_args, _info) do
    {:ok, Repo.all(Post)}
  end
end