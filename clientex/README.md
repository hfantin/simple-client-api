# Clientex

To start your Phoenix server:

  * Install dependencies with `mix deps.get`
  * Create and migrate your database with `mix ecto.setup`
  * Start Phoenix endpoint with `mix phx.server`

Now you can visit [`localhost:4000`](http://localhost:4000) from your browser.

Ready to run in production? Please [check our deployment guides](https://hexdocs.pm/phoenix/deployment.html).

## Learn more

  * Official website: http://www.phoenixframework.org/
  * Guides: https://hexdocs.pm/phoenix/overview.html
  * Docs: https://hexdocs.pm/phoenix
  * Mailing list: http://groups.google.com/group/phoenix-talk
  * Source: https://github.com/phoenixframework/phoenix


# Tutorial

- install [phoenix](https://hexdocs.pm/phoenix/up_and_running.html):   
- mix local.hex   
- mix archive.install hex phx_new 1.4.10   
- mix phx.new clientex --no-html --no-webpack --database mysql 
- cd clientex
- add  user, pass and database to config/dev.exs and  config/test.exs
- execute mix ecto.create - to drop db mix ecto.drop
- mix phx.gen 
- mix ecto.load or  mix ecto.create  -- to create database
- mix phx.gen.context CliCtx Client cli id:integer:unique name:string birth_date:date email:string
- mix phx.gen.json CliCtx Client cli id:integer name:string birth_date:date email:string --no-context --no-schema


# Run 
- mix deps.get
- mix phx.server


# Run production
[deployment](https://hexdocs.pm/phoenix/deployment.html)


# Links
[elixir performance](https://medium.com/@ron.arts/optimizing-elixir-phoenix-performance-a50f7c92b9e4)