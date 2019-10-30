defmodule Clientex2.Repo do
  use Ecto.Repo,
    otp_app: :clientex2,
    adapter: Ecto.Adapters.MyXQL
end
