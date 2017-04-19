defmodule ChinguCentral.Repo.Migrations.CreateChinguCentral.Applications.Cohort do
  use Ecto.Migration

  def change do
    create table(:applications_cohorts) do
      add :name, :string, null: false
      add :description, :text, null: false

      timestamps()
    end

    create unique_index(:applications_cohorts, [:name])
  end
end
