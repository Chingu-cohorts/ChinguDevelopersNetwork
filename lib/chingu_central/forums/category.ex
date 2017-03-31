defmodule ChinguCentral.Forums.Category do
  use Ecto.Schema

  schema "forums_categories" do
    field :description, :string
    field :name, :string

    timestamps()
  end
end
