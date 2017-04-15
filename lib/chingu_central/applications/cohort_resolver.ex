defmodule ChinguCentral.Applications.CohortResolver do
  alias ChinguCentral.Repo
  alias ChinguCentral.Applications.Cohort
  
  def all(_args, _info) do
    {:ok, Repo.all(Cohort)}
  end
  
  def find(%{id: id}, _info) do
    case Repo.get(Cohort, id) do
      nil -> {:error, "Cohort with ID #{id} not found"}
      cohort -> {:ok, cohort}
    end
  end
end