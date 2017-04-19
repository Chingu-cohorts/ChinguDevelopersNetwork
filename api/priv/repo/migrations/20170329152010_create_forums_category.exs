defmodule ChinguCentral.Repo.Migrations.CreateChinguCentral.Forums.Category do
  use Ecto.Migration

  def change do
    create table(:forums_categories) do
      add :name, :string, null: false
      add :description, :text, null: false

      timestamps()
    end

    create unique_index(:forums_categories, [:name])
  end
end
