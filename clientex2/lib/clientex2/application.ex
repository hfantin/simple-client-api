defmodule Clientex2.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc "OTP Application specification for Clientex2 - simple client api-"

  use Application
  require Logger

  def start(_type, _args) do
    # List all child processes to be supervised
    # children = [
    #   {Plug.Cowboy, scheme: :http, plug: Clientex2.Router, options: [port: 3000]}
    # ]
    port = Application.get_env(:clientex2, :port)
    children = [
      # Use Plug.Cowboy.child_spec/3 to register our endpoint as a plug
      Plug.Cowboy.child_spec(
        scheme: :http,
        plug: Clientex2.Router,
        options: [port: port]
      ), 
      Clientex2.Repo
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Clientex2.Supervisor]

    Logger.info("Starting application on port #{port}")
    Supervisor.start_link(children, opts)
  end
end
