defmodule Clientex.Repo do
  use Ecto.Repo,
    otp_app: :clientex,
    adapter: Ecto.Adapters.MyXQL
end
