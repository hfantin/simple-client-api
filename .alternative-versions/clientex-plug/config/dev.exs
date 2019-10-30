use Mix.Config

config :clientex2, port: 4001

config :logger, level: :info

config :clientex2, Clientex2.Repo,
  username: "guest",
  password: "guest",
  database: "test",
  hostname: "localhost",
  show_sensitive_data_on_connection_error: true,
  pool_size: 20