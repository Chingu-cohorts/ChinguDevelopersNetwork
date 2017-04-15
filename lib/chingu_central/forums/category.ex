defmodule ChinguCentral.Forums.Category do
  use Ecto.Schema

  schema "forums_categories" do
    field :description, :string
    field :name, :string

    has_many :forums_posts, ChinguCentral.Forums.Post, foreign_key: :forums_categories_id

    timestamps()
  end
end
