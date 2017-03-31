defmodule ChinguCentral.AccountsTest do
  use ChinguCentral.DataCase

  alias ChinguCentral.Accounts
  alias ChinguCentral.Accounts.User

  @create_attrs %{email: "some email", encrypted_password: "some encrypted_password", name: "some name", username: "some username"}
  @update_attrs %{email: "some updated email", encrypted_password: "some updated encrypted_password", name: "some updated name", username: "some updated username"}
  @invalid_attrs %{email: nil, encrypted_password: nil, name: nil, username: nil}

  def fixture(:user, attrs \\ @create_attrs) do
    {:ok, user} = Accounts.create_user(attrs)
    user
  end

  test "list_users/1 returns all users" do
    user = fixture(:user)
    assert Accounts.list_users() == [user]
  end

  test "get_user! returns the user with given id" do
    user = fixture(:user)
    assert Accounts.get_user!(user.id) == user
  end

  test "create_user/1 with valid data creates a user" do
    assert {:ok, %User{} = user} = Accounts.create_user(@create_attrs)
    assert user.email == "some email"
    assert user.encrypted_password == "some encrypted_password"
    assert user.name == "some name"
    assert user.username == "some username"
  end

  test "create_user/1 with invalid data returns error changeset" do
    assert {:error, %Ecto.Changeset{}} = Accounts.create_user(@invalid_attrs)
  end

  test "update_user/2 with valid data updates the user" do
    user = fixture(:user)
    assert {:ok, user} = Accounts.update_user(user, @update_attrs)
    assert %User{} = user
    assert user.email == "some updated email"
    assert user.encrypted_password == "some updated encrypted_password"
    assert user.name == "some updated name"
    assert user.username == "some updated username"
  end

  test "update_user/2 with invalid data returns error changeset" do
    user = fixture(:user)
    assert {:error, %Ecto.Changeset{}} = Accounts.update_user(user, @invalid_attrs)
    assert user == Accounts.get_user!(user.id)
  end

  test "delete_user/1 deletes the user" do
    user = fixture(:user)
    assert {:ok, %User{}} = Accounts.delete_user(user)
    assert_raise Ecto.NoResultsError, fn -> Accounts.get_user!(user.id) end
  end

  test "change_user/1 returns a user changeset" do
    user = fixture(:user)
    assert %Ecto.Changeset{} = Accounts.change_user(user)
  end
end
