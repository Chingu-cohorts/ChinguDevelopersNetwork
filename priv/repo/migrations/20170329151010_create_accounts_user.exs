defmodule ChinguCentral.Repo.Migrations.CreateChinguCentral.Accounts.User do
  use Ecto.Migration

  def change do
    create table(:accounts_users) do
      add :username, :string, null: false
      add :email, :string, null: false
      add :name, :string, null: false
      add :encrypted_password, :string

      timestamps()
    end

    create unique_index(:accounts_users, [:username])
    create unique_index(:accounts_users, [:email])
  end
end
