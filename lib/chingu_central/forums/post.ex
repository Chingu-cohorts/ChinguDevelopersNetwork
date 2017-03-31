defmodule ChinguCentral.Forums.Post do
  use Ecto.Schema

  schema "forums_posts" do
    field :content, :string
    field :title, :string
    # field :accounts_users_id, :id
    # field :forums_categories_id, :id
    
    belongs_to :accounts_users, ChinguCentral.Accounts.User
    belongs_to :forums_categories, ChinguCentral.Forums.Category

    timestamps()
  end
end
