defmodule ChinguCentral.Accounts.UserResolver do
  alias ChinguCentral.Repo
  alias ChinguCentral.Accounts.User
  
  def all(_args, _info) do
    {:ok, Repo.all(User)}
  end
end
