defmodule ChinguCentral.Web.Router do
  use ChinguCentral.Web, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api", ChinguCentral.Web do
    pipe_through :api
  end
  
  forward "/api", Absinthe.Plug,
    schema: ChinguCentral.Schema

  forward "/graphiql", Absinthe.Plug.GraphiQL,
    schema: ChinguCentral.Schema
end
