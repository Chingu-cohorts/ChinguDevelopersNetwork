defmodule ChinguCentral.Accounts.User do
  use Ecto.Schema

  schema "accounts_users" do
    field :email, :string
    field :encrypted_password, :string
    field :name, :string
    field :username, :string
    
    has_many :forums_posts, ChinguCentral.Forums.Post
    many_to_many :applications_cohorts, ChinguCentral.Applications.Cohort, join_through: "applications_user_cohorts"

    timestamps()
  end
end
