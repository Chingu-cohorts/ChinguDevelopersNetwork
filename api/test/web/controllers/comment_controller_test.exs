defmodule ChinguCentral.Web.CommentControllerTest do
  use ChinguCentral.Web.ConnCase

  alias ChinguCentral.Forums
  alias ChinguCentral.Forums.Comment

  @create_attrs %{content: "some content"}
  @update_attrs %{content: "some updated content"}
  @invalid_attrs %{content: nil}

  def fixture(:comment) do
    {:ok, comment} = Forums.create_comment(@create_attrs)
    comment
  end

  setup %{conn: conn} do
    {:ok, conn: put_req_header(conn, "accept", "application/json")}
  end

  test "lists all entries on index", %{conn: conn} do
    conn = get conn, comment_path(conn, :index)
    assert json_response(conn, 200)["data"] == []
  end

  test "creates comment and renders comment when data is valid", %{conn: conn} do
    conn = post conn, comment_path(conn, :create), comment: @create_attrs
    assert %{"id" => id} = json_response(conn, 201)["data"]

    conn = get conn, comment_path(conn, :show, id)
    assert json_response(conn, 200)["data"] == %{
      "id" => id,
      "content" => "some content"}
  end

  test "does not create comment and renders errors when data is invalid", %{conn: conn} do
    conn = post conn, comment_path(conn, :create), comment: @invalid_attrs
    assert json_response(conn, 422)["errors"] != %{}
  end

  test "updates chosen comment and renders comment when data is valid", %{conn: conn} do
    %Comment{id: id} = comment = fixture(:comment)
    conn = put conn, comment_path(conn, :update, comment), comment: @update_attrs
    assert %{"id" => ^id} = json_response(conn, 200)["data"]

    conn = get conn, comment_path(conn, :show, id)
    assert json_response(conn, 200)["data"] == %{
      "id" => id,
      "content" => "some updated content"}
  end

  test "does not update chosen comment and renders errors when data is invalid", %{conn: conn} do
    comment = fixture(:comment)
    conn = put conn, comment_path(conn, :update, comment), comment: @invalid_attrs
    assert json_response(conn, 422)["errors"] != %{}
  end

  test "deletes chosen comment", %{conn: conn} do
    comment = fixture(:comment)
    conn = delete conn, comment_path(conn, :delete, comment)
    assert response(conn, 204)
    assert_error_sent 404, fn ->
      get conn, comment_path(conn, :show, comment)
    end
  end
end
