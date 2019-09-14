defmodule Clientex.CliCtx.Client do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :id, autogenerate: true}
  schema "CLI" do
    # field :id, :integer
    field :name, :string
    field :birth_date, :date
    field :email, :string
    

    # timestamps()
  end

  @doc false
  def changeset(client, attrs) do
    client
    |> cast(attrs, [:id, :name, :birth_date, :email])
    |> validate_required([:id, :name, :birth_date, :email])
    |> unique_constraint(:id)
  end
end
