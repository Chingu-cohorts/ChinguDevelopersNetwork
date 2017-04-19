defmodule ChinguCentral.Repo.Migrations.CreateChinguCentral.Forums.Comment do
  use Ecto.Migration

  def change do
    create table(:forums_comments) do
      add :content, :text, null: false
      add :accounts_users_id, references(:accounts_users, on_delete: :nothing)
      add :forums_posts_id, references(:forums_posts, on_delete: :nothing)

      timestamps()
    end

    create index(:forums_comments, [:accounts_users_id])
    create index(:forums_comments, [:forums_posts_id])
  end
end
