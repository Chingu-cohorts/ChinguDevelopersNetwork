defmodule ChinguCentral.Forums.Post do
  use Ecto.Schema

  schema "forums_posts" do
    field :content, :string
    field :title, :string
    # field :accounts_users_id, :id
    # field :forums_categories_id, :id
    
    belongs_to :accounts_user, ChinguCentral.Accounts.User, foreign_key: :accounts_users_id
    belongs_to :forums_category, ChinguCentral.Forums.Category, foreign_key: :forums_categories_id

    timestamps()
  end
end
