defmodule ChinguCentral.Applications.UserCohort do
  use Ecto.Schema

  schema "applications_user_cohorts" do
    belongs_to :accounts_users, ChinguCentral.Accounts.User, foreign_key: :accounts_users_id
    belongs_to :applications_cohorts, ChinguCentral.Applications.Cohort, foreign_key: :applications_cohorts_id

    timestamps()
  end
end