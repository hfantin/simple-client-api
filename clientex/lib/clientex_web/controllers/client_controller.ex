defmodule ClientexWeb.ClientController do
  use ClientexWeb, :controller

  alias Clientex.CliCtx
  alias Clientex.CliCtx.Client

  action_fallback ClientexWeb.FallbackController

  def index(conn, _params) do
    cli = CliCtx.list_cli()
    render(conn, "index.json", cli: cli)
  end

  def create(conn, %{"client" => client_params}) do
    with {:ok, %Client{} = client} <- CliCtx.create_client(client_params) do
      conn
      |> put_status(:created)
      |> put_resp_header("location", Routes.client_path(conn, :show, client))
      |> render("show.json", client: client)
    end
  end

  def show(conn, %{"id" => id}) do
    client = CliCtx.get_client!(id)
    render(conn, "show.json", client: client)
  end

  def update(conn, %{"id" => id, "client" => client_params}) do
    client = CliCtx.get_client!(id)

    with {:ok, %Client{} = client} <- CliCtx.update_client(client, client_params) do
      render(conn, "show.json", client: client)
    end
  end

  def delete(conn, %{"id" => id}) do
    client = CliCtx.get_client!(id)

    with {:ok, %Client{}} <- CliCtx.delete_client(client) do
      send_resp(conn, :no_content, "")
    end
  end
end
