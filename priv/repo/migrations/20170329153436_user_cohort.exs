defmodule ChinguCentral.Repo.Migrations.UserCohort do
  use Ecto.Migration

  def change do
    create table(:user_cohorts) do
      add :user_id, references(:accounts_users, on_delete: :nothing), null: false
      add :cohort_id, references(:applications_cohorts, on_delete: :nothing), null: false

      timestamps()
    end
    create index(:user_cohorts, [:user_id])
    create index(:user_cohorts, [:cohort_id])
    create index(:user_cohorts, [:user_id, :cohort_id], unique: true)
  end
end
