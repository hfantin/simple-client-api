defmodule Clientex.Repo.Migrations.CreateCli do
  use Ecto.Migration

  def change do
    create table(:cli, primary_key: false) do
      # add :id, :integer
      add :id, :serial, primary_key: true
      add :name, :string
      add :birth_date, :date
      add :email, :string, null: true

      timestamps()
    end

    # create unique_index(:cli, [:id])
  end
end
