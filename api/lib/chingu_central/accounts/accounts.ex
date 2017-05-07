defmodule ChinguCentral.Accounts do
  @moduledoc """
  The boundary for the Accounts system.
  """

  import Ecto.{Query, Changeset}, warn: false

  alias ChinguCentral.Accounts.User

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking user changes.

  ## Examples

      iex> change_user(user)
      %Ecto.Changeset{source: %User{}}

  """
  def change_user(%User{} = user) do
    user_changeset(user, %{})
  end

  defp user_changeset(%User{} = user, attrs) do
    user
    |> cast(attrs, [:username, :email, :name, :encrypted_password])
    |> validate_required([:username, :email, :name, :encrypted_password])
    |> validate_format(:email, ~r/[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z]+/)
    |> validate_length(:email, min: 6, max: 40)
    |> validate_length(:username, min: 3, max: 30)
    |> unique_constraint(:username)
    |> unique_constraint(:email)
  end
end
