defmodule ChinguCentral.Applications do
  @moduledoc """
  The boundary for the Applications system.
  """

  import Ecto.{Query, Changeset}, warn: false

  alias ChinguCentral.Applications.{Cohort, UserCohort}

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking cohort changes.

  ## Examples

      iex> change_cohort(cohort)
      %Ecto.Changeset{source: %Cohort{}}

  """
  def change_cohort(%Cohort{} = cohort) do
    cohort_changeset(cohort, %{})
  end

  defp cohort_changeset(%Cohort{} = cohort, attrs) do
    cohort
    |> cast(attrs, [:name, :description])
    |> validate_required([:name, :description])
    |> unique_constraint(:name)       
    |> validate_length(:name, min:3, max:30)
    |> validate_format(:name, ~r/^[a-zA-Z]+$/)
  end
  
  @doc """
  Returns an `%Ecto.Changeset{}` for tracking cohort changes.

  ## Examples

      iex> change_user_cohort(user_cohort)
      %Ecto.Changeset{source: %UserCohort{}}

  """
  def change_user_cohort(%UserCohort{} = user_cohort) do
    user_cohort_changeset(user_cohort, %{})
  end
  
  defp user_cohort_changeset(%UserCohort{} = user_cohort, attrs) do
    user_cohort
    |> cast(attrs, [:user_id, :cohort_id])
    |> validate_required([:user_id, :cohort_id])
    |> unique_constraint(:user_id_cohort_id)
  end
end
