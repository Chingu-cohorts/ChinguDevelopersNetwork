defmodule ChinguCentral.ApplicationsTest do
  use ChinguCentral.DataCase

  alias ChinguCentral.Applications
  alias ChinguCentral.Applications.Cohort

  @create_attrs %{description: "some description", name: "some name"}
  @update_attrs %{description: "some updated description", name: "some updated name"}
  @invalid_attrs %{description: nil, name: nil}

  def fixture(:cohort, attrs \\ @create_attrs) do
    {:ok, cohort} = Applications.create_cohort(attrs)
    cohort
  end

"""
  test "list_cohorts/1 returns all cohorts" do
    cohort = fixture(:cohort)
    assert Applications.list_cohorts() == [cohort]
  end

  test "get_cohort! returns the cohort with given id" do
    cohort = fixture(:cohort)
    assert Applications.get_cohort!(cohort.id) == cohort
  end

  test "create_cohort/1 with valid data creates a cohort" do
    assert {:ok, %Cohort{} = cohort} = Applications.create_cohort(@create_attrs)
    assert cohort.description == "some description"
    assert cohort.name == "some name"
  end

  test "create_cohort/1 with invalid data returns error changeset" do
    assert {:error, %Ecto.Changeset{}} = Applications.create_cohort(@invalid_attrs)
  end

  test "update_cohort/2 with valid data updates the cohort" do
    cohort = fixture(:cohort)
    assert {:ok, cohort} = Applications.update_cohort(cohort, @update_attrs)
    assert %Cohort{} = cohort
    assert cohort.description == "some updated description"
    assert cohort.name == "some updated name"
  end

  test "update_cohort/2 with invalid data returns error changeset" do
    cohort = fixture(:cohort)
    assert {:error, %Ecto.Changeset{}} = Applications.update_cohort(cohort, @invalid_attrs)
    assert cohort == Applications.get_cohort!(cohort.id)
  end

  test "delete_cohort/1 deletes the cohort" do
    cohort = fixture(:cohort)
    assert {:ok, %Cohort{}} = Applications.delete_cohort(cohort)
    assert_raise Ecto.NoResultsError, fn -> Applications.get_cohort!(cohort.id) end
  end

  test "change_cohort/1 returns a cohort changeset" do
    cohort = fixture(:cohort)
    assert %Ecto.Changeset{} = Applications.change_cohort(cohort)
  end
"""
end
