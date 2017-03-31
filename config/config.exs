# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.
use Mix.Config

# General application configuration
config :chingu_central,
  ecto_repos: [ChinguCentral.Repo]

# Configures the endpoint
config :chingu_central, ChinguCentral.Web.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "W6xnoPooHXsT0fqupSHLEHATHiPZ6bwY9uARuurXmFWPdyxpXHGpYYQbQe4O+duL",
  render_errors: [view: ChinguCentral.Web.ErrorView, accepts: ~w(json)],
  pubsub: [name: ChinguCentral.PubSub,
           adapter: Phoenix.PubSub.PG2]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env}.exs"
