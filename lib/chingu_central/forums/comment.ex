defmodule ChinguCentral.Forums.Comment do
  use Ecto.Schema

  schema "forums_comments" do
    field :content, :string
    
    belongs_to :accounts_users, ChinguCentral.Accounts.User
    belongs_to :forums_posts, ChinguCentral.Forums.Post
    # field :user_id, :id
    # field :post_id, :id

    timestamps()
  end
end
