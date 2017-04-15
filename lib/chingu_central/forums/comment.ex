defmodule ChinguCentral.Forums.Comment do
  use Ecto.Schema

  schema "forums_comments" do
    field :content, :string
    
    belongs_to :accounts_users, ChinguCentral.Accounts.User, foreign_key: :accounts_users_id
    belongs_to :forums_posts, ChinguCentral.Forums.Post, foreign_key: :forums_posts_id

    timestamps()
  end
end
