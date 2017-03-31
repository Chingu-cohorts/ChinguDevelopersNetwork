defmodule ChinguCentral.Repo.Migrations.CreateChinguCentral.Forums.Post do
  use Ecto.Migration

  def change do
    create table(:forums_posts) do
      add :title, :string, null: false
      add :content, :text, null: false
      add :accounts_users_id, references(:accounts_users, on_delete: :nothing)
      add :forums_categories_id, references(:forums_categories, on_delete: :nothing)

      timestamps()
    end

    create index(:forums_posts, [:accounts_users_id])
    create index(:forums_posts, [:forums_categories_id])
  end
end
