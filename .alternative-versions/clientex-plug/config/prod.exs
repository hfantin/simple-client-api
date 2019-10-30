use Mix.Config

config :clientex2, port: 4000

config :logger, level: :error

config :clientex2, Clientex2.Repo,
  database: "test",
  username: "guest",
  password: "guest",
  hostname: "localhost",
  show_sensitive_data_on_connection_error: true,
  pool_size: 20