defmodule ChinguCentral.Forums.Post do
  use Ecto.Schema

  schema "forums_posts" do
    field :content, :string
    field :title, :string
    
    belongs_to :accounts_user, ChinguCentral.Accounts.User, foreign_key: :accounts_users_id
    belongs_to :forums_category, ChinguCentral.Forums.Category, foreign_key: :forums_categories_id
    
    has_many :forums_comments, ChinguCentral.Forums.Comment, foreign_key: :forums_posts_id

    timestamps()
  end
end
