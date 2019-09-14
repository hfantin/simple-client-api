defmodule ClientexWeb.Router do
  use ClientexWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api", ClientexWeb do
    pipe_through :api
    resources "/clients", ClientController, except: [:new, :edit]
  end
end
