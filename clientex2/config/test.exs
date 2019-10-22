use Mix.Config

config :clientex2, port: 4002

config :clientex2, Clientex2.Repo,
  username: "guest",
  password: "guest",
  database: "test",
  hostname: "localhost",
  pool: Ecto.Adapters.SQL.Sandbox
