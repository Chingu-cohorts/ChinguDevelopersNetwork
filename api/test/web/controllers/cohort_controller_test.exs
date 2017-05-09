defmodule ChinguCentral.Web.CohortControllerTest do
  use ChinguCentral.Web.ConnCase

  alias ChinguCentral.Applications
  alias ChinguCentral.Applications.Cohort

  @create_attrs %{description: "some description", name: "some name"}
  @update_attrs %{description: "some updated description", name: "some updated name"}
  @invalid_attrs %{description: nil, name: nil}

  def fixture(:cohort) do
    {:ok, cohort} = Applications.create_cohort(@create_attrs)
    cohort
  end

  setup %{conn: conn} do
    {:ok, conn: put_req_header(conn, "accept", "application/json")}
  end

"""
  test "lists all entries on index", %{conn: conn} do
    conn = get conn, cohort_path(conn, :index)
    assert json_response(conn, 200)["data"] == []
  end

  test "creates cohort and renders cohort when data is valid", %{conn: conn} do
    conn = post conn, cohort_path(conn, :create), cohort: @create_attrs
    assert %{"id" => id} = json_response(conn, 201)["data"]

    conn = get conn, cohort_path(conn, :show, id)
    assert json_response(conn, 200)["data"] == %{
      "id" => id,
      "description" => "some description",
      "name" => "some name"}
  end

  test "does not create cohort and renders errors when data is invalid", %{conn: conn} do
    conn = post conn, cohort_path(conn, :create), cohort: @invalid_attrs
    assert json_response(conn, 422)["errors"] != %{}
  end

  test "updates chosen cohort and renders cohort when data is valid", %{conn: conn} do
    %Cohort{id: id} = cohort = fixture(:cohort)
    conn = put conn, cohort_path(conn, :update, cohort), cohort: @update_attrs
    assert %{"id" => ^id} = json_response(conn, 200)["data"]

    conn = get conn, cohort_path(conn, :show, id)
    assert json_response(conn, 200)["data"] == %{
      "id" => id,
      "description" => "some updated description",
      "name" => "some updated name"}
  end

  test "does not update chosen cohort and renders errors when data is invalid", %{conn: conn} do
    cohort = fixture(:cohort)
    conn = put conn, cohort_path(conn, :update, cohort), cohort: @invalid_attrs
    assert json_response(conn, 422)["errors"] != %{}
  end

  test "deletes chosen cohort", %{conn: conn} do
    cohort = fixture(:cohort)
    conn = delete conn, cohort_path(conn, :delete, cohort)
    assert response(conn, 204)
    assert_error_sent 404, fn ->
      get conn, cohort_path(conn, :show, cohort)
    end
  end
"""
end
