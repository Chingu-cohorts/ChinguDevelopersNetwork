defmodule ChinguCentral.Forums.CategoryResolver do
  alias ChinguCentral.Repo
  alias ChinguCentral.Forums.Category
  
  def all(_args, _info) do
    {:ok, Repo.all(Category)}
  end
  
  def find(%{id: id}, _info) do
    case Repo.get(Category, id) do
      nil -> {:error, "Category with ID #{id} not found"}
      category -> {:ok, category}
    end
  end
end