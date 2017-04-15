defmodule ChinguCentral.Forums.CategoryResolver do
  alias ChinguCentral.Repo
  alias ChinguCentral.Forums.Category
  
  def all(_args, _info) do
    {:ok, Repo.all(Category)}
  end
end