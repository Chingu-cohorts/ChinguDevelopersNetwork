defmodule ChinguCentral.Repo.Migrations.UpdateUserAccountsTable do
  use Ecto.Migration

  def change do
    alter table(:accounts_users) do
      add :hashed_password, :string
      remove :encrypted_password
    end
  end
end
