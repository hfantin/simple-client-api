defmodule Clientex.CliCtxTest do
  use Clientex.DataCase

  alias Clientex.CliCtx

  describe "cli" do
    alias Clientex.CliCtx.Client

    @valid_attrs %{birth_date: ~D[2010-04-17], email: "some email", id: 42, name: "some name"}
    @update_attrs %{birth_date: ~D[2011-05-18], email: "some updated email", id: 43, name: "some updated name"}
    @invalid_attrs %{birth_date: nil, email: nil, id: nil, name: nil}

    def client_fixture(attrs \\ %{}) do
      {:ok, client} =
        attrs
        |> Enum.into(@valid_attrs)
        |> CliCtx.create_client()

      client
    end

    test "list_cli/0 returns all cli" do
      client = client_fixture()
      assert CliCtx.list_cli() == [client]
    end

    test "get_client!/1 returns the client with given id" do
      client = client_fixture()
      assert CliCtx.get_client!(client.id) == client
    end

    test "create_client/1 with valid data creates a client" do
      assert {:ok, %Client{} = client} = CliCtx.create_client(@valid_attrs)
      assert client.birth_date == ~D[2010-04-17]
      assert client.email == "some email"
      assert client.id == 42
      assert client.name == "some name"
    end

    test "create_client/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = CliCtx.create_client(@invalid_attrs)
    end

    test "update_client/2 with valid data updates the client" do
      client = client_fixture()
      assert {:ok, %Client{} = client} = CliCtx.update_client(client, @update_attrs)
      assert client.birth_date == ~D[2011-05-18]
      assert client.email == "some updated email"
      assert client.id == 43
      assert client.name == "some updated name"
    end

    test "update_client/2 with invalid data returns error changeset" do
      client = client_fixture()
      assert {:error, %Ecto.Changeset{}} = CliCtx.update_client(client, @invalid_attrs)
      assert client == CliCtx.get_client!(client.id)
    end

    test "delete_client/1 deletes the client" do
      client = client_fixture()
      assert {:ok, %Client{}} = CliCtx.delete_client(client)
      assert_raise Ecto.NoResultsError, fn -> CliCtx.get_client!(client.id) end
    end

    test "change_client/1 returns a client changeset" do
      client = client_fixture()
      assert %Ecto.Changeset{} = CliCtx.change_client(client)
    end
  end
end
