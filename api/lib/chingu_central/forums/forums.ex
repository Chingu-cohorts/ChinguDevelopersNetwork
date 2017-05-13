defmodule ChinguCentral.Forums do
  @moduledoc """
  The boundary for the Forums system.
  """

  import Ecto.{Query, Changeset}, warn: false
  alias ChinguCentral.Repo

  alias ChinguCentral.Forums.{Category, Post, Comment}

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking category changes.

  ## Examples

      iex> change_category(category)
      %Ecto.Changeset{source: %Category{}}

  """
  def change_category(%Category{} = category) do
    category_changeset(category, %{})
  end

  defp category_changeset(%Category{} = category, attrs) do
    category
    |> cast(attrs, [:name, :description])
    |> validate_required([:name, :description])
    |> unique_constraint(:name)
    |> validate_length(:name, min: 6)
    |> validate_length(:name, max: 30)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking post changes.

  ## Examples

      iex> change_post(post)
      %Ecto.Changeset{source: %Post{}}

  """
  def change_post(%Post{} = post) do
    post_changeset(post, %{})
  end

  defp post_changeset(%Post{} = post, attrs) do
    post
    |> cast(attrs, [:title, :content, :forums_categories_id, :accounts_users_id])
    |> validate_required([:title, :content, :forums_categories_id, :accounts_users_id])
  end 

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking comment changes.

  ## Examples

      iex> change_comment(comment)
      %Ecto.Changeset{source: %Comment{}}

  """
  def change_comment(%Comment{} = comment) do
    comment_changeset(comment, %{})
  end

  defp comment_changeset(%Comment{} = comment, attrs) do
    comment
    |> cast(attrs, [:content, :forums_posts_id, :accounts_users_id])
    |> validate_required([:content, :forums_posts_id, :accounts_users_id])
  end
end
