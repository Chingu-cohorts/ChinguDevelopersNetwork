defmodule ChinguCentral.Schema do
  use Absinthe.Schema
  import_types ChinguCentral.Schema.Types
  
  query do
    field :accounts_users, list_of(:accounts_user) do
      resolve &ChinguCentral.Accounts.UserResolver.all/2
    end
    
    field :accounts_user, type: :accounts_user do
      arg :id, non_null(:id)
      resolve &ChinguCentral.Accounts.UserResolver.find/2
    end
    
    field :forums_posts, list_of(:forums_post) do
      resolve &ChinguCentral.Forums.PostResolver.all/2
    end
    
    field :forums_post, type: :forums_post do
      arg :id, non_null(:id)
      resolve &ChinguCentral.Forums.PostResolver.find/2
    end
    
    field :forums_categories, list_of(:forums_category) do
      resolve &ChinguCentral.Forums.CategoryResolver.all/2
    end
    
    field :forums_category, type: :forums_category do
      arg :id, non_null(:id)
      resolve &ChinguCentral.Forums.CategoryResolver.find/2
    end
    
    field :forums_comments, list_of(:forums_comment) do
      resolve &ChinguCentral.Forums.CommentResolver.all/2
    end
    
    field :forums_comment, type: :forums_comment do
      arg :id, non_null(:id)
      resolve &ChinguCentral.Forums.CommentResolver.find/2
    end
    
    field :applications_cohorts, list_of(:applications_cohort) do
      resolve &ChinguCentral.Applications.CohortResolver.all/2
    end
    
    field :applications_cohort, type: :applications_cohort do
      arg :id, non_null(:id)
      resolve &ChinguCentral.Applications.CohortResolver.find/2
    end
  end
end