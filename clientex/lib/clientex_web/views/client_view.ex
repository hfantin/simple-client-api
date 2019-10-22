defmodule ClientexWeb.ClientView do
  use ClientexWeb, :view
  alias ClientexWeb.ClientView

  def render("index.json", %{cli: cli}) do
    %{data: render_many(cli, ClientView, "client.json")}
  end

  def render("show.json", %{client: client}) do
    %{data: render_one(client, ClientView, "client.json")}
  end

  def render("client.json", %{client: client}) do
    %{id: client.id,
      name: client.name,
      birth_date: client.birth_date,
      email: client.email}
  end
end
