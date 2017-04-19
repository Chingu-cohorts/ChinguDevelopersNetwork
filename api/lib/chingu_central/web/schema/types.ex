defmodule ChinguCentral.Schema.Types do
  use Absinthe.Schema.Notation
  use Absinthe.Ecto, repo: ChinguCentral.Repo
  
  object :accounts_user do
    field :id, :id
    field :name, :string
    field :email, :string
    field :username, :string
    field :forums_posts, list_of(:forums_post), resolve: assoc(:forums_posts)
    field :forums_comments, list_of(:forums_comment), resolve: assoc(:forums_comments)
    field :applications_cohorts, list_of(:applications_cohort), resolve: assoc(:applications_cohorts)
  end
  
  object :applications_cohort do
    field :id, :id
    field :name, :string
    field :description, :string
    field :accounts_users, list_of(:accounts_user), resolve: assoc(:accounts_users)
  end
  
  object :forums_post do
    field :id, :id
    field :title, :string
    field :content, :string
    field :accounts_user, :accounts_user, resolve: assoc(:accounts_user)
    field :forums_category, :forums_category, resolve: assoc(:forums_category)
    field :forums_comments, list_of(:forums_comment), resolve: assoc(:forums_comments)
  end
  
  object :forums_comment do
    field :id, :id
    field :content, :string
    field :accounts_user, :accounts_user, resolve: assoc(:accounts_users)
    field :forums_post, :forums_post, resolve: assoc(:forums_posts)
  end
  
  object :forums_category do
    field :id, :id
    field :name, :string
    field :description, :string
    field :forums_posts, list_of(:forums_post), resolve: assoc(:forums_posts)
  end
end