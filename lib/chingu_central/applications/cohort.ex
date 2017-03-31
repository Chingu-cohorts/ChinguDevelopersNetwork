defmodule ChinguCentral.Applications.Cohort do
  use Ecto.Schema

  schema "applications_cohorts" do
    field :description, :string
    field :name, :string

    many_to_many :accounts_users, ChinguCentral.Accounts.User, join_through: "applications_user_cohorts"

    timestamps()
  end
end
