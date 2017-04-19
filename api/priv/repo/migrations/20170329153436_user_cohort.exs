defmodule ChinguCentral.Repo.Migrations.UserCohort do
  use Ecto.Migration

  def change do
    create table(:applications_user_cohorts) do
      add :accounts_users_id, references(:accounts_users, on_delete: :nothing), null: false
      add :applications_cohorts_id, references(:applications_cohorts, on_delete: :nothing), null: false

      timestamps()
    end
    create index(:applications_user_cohorts, [:accounts_users_id])
    create index(:applications_user_cohorts, [:applications_cohorts_id])
    create index(:applications_user_cohorts, [:accounts_users_id, :applications_cohorts_id], unique: true)
  end
end
