defmodule Clientex2.Router do
  use Plug.Router
  require Logger
  alias Clientex2.Plug.VerifyRequest

  plug Plug.Parsers,
    parsers: [:json],
    pass: ["application/json"],
    json_decoder: Poison
  plug :match
  plug :dispatch

  get "/v1/clients" do
    # Logger.info("router /v1/clients")
    Clientex2.Api.Clients.get_clients(conn)
  end

  match _ do
    send_resp(conn, 404, "pagina não existe")
  end
end