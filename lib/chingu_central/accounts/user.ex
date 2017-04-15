defmodule ChinguCentral.Accounts.User do
  use Ecto.Schema

  schema "accounts_users" do
    field :email, :string
    field :encrypted_password, :string
    field :name, :string
    field :username, :string
    
    has_many :forums_posts, ChinguCentral.Forums.Post, foreign_key: :accounts_users_id
    many_to_many :applications_cohorts, ChinguCentral.Applications.Cohort, join_through: "applications_user_cohorts", join_keys: [accounts_users_id: :id, applications_cohorts_id: :id]

    timestamps()
  end
end
