defmodule ChinguCentral.Applications.UserCohort do
  use Ecto.Schema

  schema "applications_user_cohorts" do
    belongs_to :accounts_users, ChinguCentral.Accounts.User
    belongs_to :applications_cohorts, ChinguCentral.Applications.Cohort

    timestamps()
  end
end