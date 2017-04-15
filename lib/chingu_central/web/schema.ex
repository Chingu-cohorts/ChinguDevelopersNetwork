defmodule ChinguCentral.Schema do
  use Absinthe.Schema
  import_types ChinguCentral.Schema.Types
  
  query do
    field :accounts_users, list_of(:accounts_user) do
      resolve &ChinguCentral.Accounts.UserResolver.all/2
    end
    
    field :forums_posts, list_of(:forums_post) do
      resolve &ChinguCentral.Forums.PostResolver.all/2
    end
    
    field :forums_categories, list_of(:forums_category) do
      resolve &ChinguCentral.Forums.CategoryResolver.all/2
    end
    
    field :applications_cohorts, list_of(:applications_cohort) do
      resolve &ChinguCentral.Applications.CohortResolver.all/2
    end
  end
end