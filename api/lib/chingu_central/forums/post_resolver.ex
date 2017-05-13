defmodule ChinguCentral.Forums.PostResolver do
  alias ChinguCentral.Repo
  alias ChinguCentral.Forums
  alias ChinguCentral.Forums.Post
  
  def all(_args, _info) do
    {:ok, Repo.all(Post)}
  end

  def create(args, _info) do
    %Post{}
    |> Forums.change_post
    |> Repo.insert
  end
  
  def find(%{id: id}, _info) do
    case Repo.get(Post, id) do
      nil -> {:error, "Post with ID #{id} not found"}
      post -> {:ok, post}
    end
  end

end