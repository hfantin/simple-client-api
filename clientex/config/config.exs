# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
use Mix.Config

config :clientex,
  ecto_repos: [Clientex.Repo]

# Configures the endpoint
config :clientex, ClientexWeb.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "0stwB7GFmmcm9CP8oZgPXya02F6VWQSxdE5qACjkpbXwiWGn9yBofYkSTQGdGRLv",
  render_errors: [view: ClientexWeb.ErrorView, accepts: ~w(json)],
  pubsub: [name: Clientex.PubSub, adapter: Phoenix.PubSub.PG2]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env()}.exs"
