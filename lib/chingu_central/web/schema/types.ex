defmodule ChinguCentral.Schema.Types do
  use Absinthe.Schema.Notation
  use Absinthe.Ecto, repo: ChinguCentral.Repo
  
  object :accounts_user do
    field :id, :id
    field :name, :string
    field :email, :string
    field :username, :string
    field :forums_posts, list_of(:forums_post), resolve: assoc(:forums_posts)
  end
  
  object :forums_post do
    field :id, :id
    field :title, :string
    field :content, :string
    field :accounts_user, :accounts_user, resolve: assoc(:accounts_user)
  end
end