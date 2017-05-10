defmodule ChinguCentral.AccountsTest do
  use ChinguCentral.DataCase

  alias ChinguCentral.Accounts
  alias ChinguCentral.Accounts.User

  @create_attrs %{email: "test1@email.com", password: "testpassword1", name: "TestUser1", username: "test1"}
  @update_attrs %{email: "test2@email.com", password: "updatetestpassword2", name: "TestUser2", username: "test2"}
  @invalid_attrs %{email: nil, encrypted_password: nil, name: nil, username: nil}

  def fixture(:user, attrs \\ @create_attrs) do
    {:ok, user} = Accounts.create_user(attrs)
    user
  end

  test "list_users/1 returns all users" do
    user = fixture(:user)
    assert Accounts.list_users() == [Map.put(user, :password, nil)]
  end

  test "get_user! returns the user with given id" do
    user = fixture(:user)
    assert Accounts.get_user!(user.id) == Map.put(user, :password, nil)
  end

  test "create_user/1 with valid data creates a user" do
    assert {:ok, %User{} = user} = Accounts.create_user(@create_attrs)
    assert user.email == @create_attrs.email
    assert user.password == @create_attrs.password
    assert Comeonin.Bcrypt.checkpw(@create_attrs.password, user.hashed_password)
    assert user.name == @create_attrs.name
    assert user.username == @create_attrs.username
  end

  test "create_user/1 with invalid data returns error changeset" do
    assert {:error, %Ecto.Changeset{}} = Accounts.create_user(@invalid_attrs)
  end

  test "update_user/2 with valid data updates the user" do
    user = fixture(:user)
    assert {:ok, user} = Accounts.update_user(user, @update_attrs)
    assert %User{} = user
    assert user.email == @update_attrs.email
    assert user.password == @update_attrs.password
    assert Comeonin.Bcrypt.checkpw(@update_attrs.password, user.hashed_password)
    assert user.name == @update_attrs.name
    assert user.username == @update_attrs.username
  end

  test "update_user/2 with invalid data returns error changeset" do
    user = fixture(:user)
    assert {:error, %Ecto.Changeset{}} = Accounts.update_user(user, @invalid_attrs)
    assert Map.put(user, :password, nil) == Accounts.get_user!(user.id)
  end

"""
  test "delete_user/1 deletes the user" do
    user = fixture(:user)
    assert {:ok, %User{}} = Accounts.delete_user(user)
    assert_raise Ecto.NoResultsError, fn -> Accounts.get_user!(user.id) end
  end

  test "change_user/1 returns a user changeset" do
    user = fixture(:user)
    assert %Ecto.Changeset{} = Accounts.change_user(user)
  end
"""
end
