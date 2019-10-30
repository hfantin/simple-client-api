defmodule Clientex2.Api.Clients do
  import Plug.Conn
  require Logger
  alias Clientex2.CliCtx
  alias Clientex2.CliCtx.Client
  
  def get_clients(conn) do
    fields = Client.__schema__(:fields)
    clients = CliCtx.list_cli()
    |> Enum.map(&Map.take(&1, fields))
    # |> Enum.map(&strip_meta/1)
    # |> Enum.into(%{})
    # |> Enum.map(&2)
    # Logger.info("get_clients")
    # Enum.each(clients, &(Logger.info("#{is_map(&1)}")))
    conn
    |> put_resp_content_type("application/json")
    |> send_resp(200, Poison.encode!(clients))
  end


  # https://medium.com/@vasspilka/serializing-ecto-models-with-poison-6df82b3e264f

  # http://blog.trenpixster.info/til-about-ecto-schema-derive/

  # def strip_meta({key, val}) when is_map(val) or is_list(val) do
  #   {key, strip_meta(val)}
  # end

  # def strip_meta(data), do: data
end