defmodule ChinguCentral.ForumsTest do
  use ChinguCentral.DataCase

  alias ChinguCentral.Forums
  alias ChinguCentral.Forums.Category

  @create_attrs %{description: "some description", name: "some name"}
  @update_attrs %{description: "some updated description", name: "some updated name"}
  @invalid_attrs %{description: nil, name: nil}

  def fixture(:category, attrs \\ @create_attrs) do
    {:ok, category} = Forums.create_category(attrs)
    category
  end

  test "list_categories/1 returns all categories" do
    category = fixture(:category)
    assert Forums.list_categories() == [category]
  end

  test "get_category! returns the category with given id" do
    category = fixture(:category)
    assert Forums.get_category!(category.id) == category
  end

  test "create_category/1 with valid data creates a category" do
    assert {:ok, %Category{} = category} = Forums.create_category(@create_attrs)
    assert category.description == "some description"
    assert category.name == "some name"
  end

  test "create_category/1 with invalid data returns error changeset" do
    assert {:error, %Ecto.Changeset{}} = Forums.create_category(@invalid_attrs)
  end

  test "update_category/2 with valid data updates the category" do
    category = fixture(:category)
    assert {:ok, category} = Forums.update_category(category, @update_attrs)
    assert %Category{} = category
    assert category.description == "some updated description"
    assert category.name == "some updated name"
  end

  test "update_category/2 with invalid data returns error changeset" do
    category = fixture(:category)
    assert {:error, %Ecto.Changeset{}} = Forums.update_category(category, @invalid_attrs)
    assert category == Forums.get_category!(category.id)
  end

  test "delete_category/1 deletes the category" do
    category = fixture(:category)
    assert {:ok, %Category{}} = Forums.delete_category(category)
    assert_raise Ecto.NoResultsError, fn -> Forums.get_category!(category.id) end
  end

  test "change_category/1 returns a category changeset" do
    category = fixture(:category)
    assert %Ecto.Changeset{} = Forums.change_category(category)
  end
end
