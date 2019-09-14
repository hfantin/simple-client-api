defmodule Clientex.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  def start(_type, _args) do
    # List all child processes to be supervised
    children = [
      # Start the Ecto repository
      Clientex.Repo,
      # Start the endpoint when the application starts
      ClientexWeb.Endpoint
      # Starts a worker by calling: Clientex.Worker.start_link(arg)
      # {Clientex.Worker, arg},
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Clientex.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  def config_change(changed, _new, removed) do
    ClientexWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
