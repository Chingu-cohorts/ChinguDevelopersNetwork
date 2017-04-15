defmodule ChinguCentral.Applications.CohortResolver do
  alias ChinguCentral.Repo
  alias ChinguCentral.Applications.Cohort
  
  def all(_args, _info) do
    {:ok, Repo.all(Cohort)}
  end
end