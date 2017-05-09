defmodule ChinguCentral.Accounts do
  @moduledoc """
  The boundary for the Accounts system.
  """

  import Ecto.{Query, Changeset}, warn: false

  alias ChinguCentral.Repo
  alias ChinguCentral.Accounts.User
  alias ChinguCentral.Accounts.UserResolver

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking user changes.

  ## Examples

      iex> change_user(user)
      %Ecto.Changeset{source: %User{}}

  """

  def list_users do
    {_, users} = UserResolver.all(%{}, %{})
    users
  end

  def get_user!(id) do
    {:ok, user} = UserResolver.find(%{id: id}, %{})
    user
  end

  def create_user(attrs) do
    changeset = get_user_changeset(%User{}, attrs)
    if changeset.valid? do
      insert_user(changeset)
    else
      {:error, changeset}
    end
  end

  def update_user(%User{} = user, attrs) do
    changeset = get_user_changeset(user, attrs)
      if changeset.valid? do
        update_user(changeset)
      else
        {:error, changeset}
    end
  end

  def hash_password_if_required(changeset) do
    case changeset do
      %Ecto.Changeset{changes: %{password: password}} ->
        put_change(changeset, :hashed_password, Comeonin.Bcrypt.hashpwsalt(password))
      _ ->
        changeset
    end
  end

  def insert_user(changeset) do
    changeset
     |> hash_password_if_required
     |> Repo.insert
  end

  def update_user(changeset) do
    changeset
     |> hash_password_if_required
     |> Repo.update
  end

  def get_user_changeset(%User{} = user, attrs) do
    user
    |> cast(attrs, [:username, :email, :name, :password])
    |> validate_required([:username, :email, :name, :password])
    |> unique_constraint(:username)
    |> unique_constraint(:email)
  end
end
