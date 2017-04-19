defmodule ChinguCentral.Applications.Cohort do
  use Ecto.Schema

  schema "applications_cohorts" do
    field :description, :string
    field :name, :string

    many_to_many :accounts_users, ChinguCentral.Accounts.User, join_through: "applications_user_cohorts", join_keys: [applications_cohorts_id: :id, accounts_users_id: :id]

    timestamps()
  end
end
