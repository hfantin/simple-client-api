# Simple client api in Elixir using Plug 
- created following [this](https://elixirschool.com/pt/lessons/specifics/plug/) tutorial 

### Steps:
1. mix new clientex2 --sup
2. add  {:plug_cowboy, "~> 2.0"}, to mix.exs
3. mix deps.get
4. mix run --no-halt
5. localhost:3000

## Dependencies:
Cowboy - A small, fast and modern HTTP server for Erlang/OTP
PLUG - 
Poison  - A JSON library for Elixir focusing on wicked-fast speed without sacrificing simplicity, completeness, or correctness.

### Links
[plug + cowboy + poison](https://blog.lelonek.me/minimal-elixir-http2-server-64188d0c1f3a)
[elixir |> cowbow |> plug](https://dev.to/jonlunsford/elixir-building-a-small-json-endpoint-with-plug-cowboy-and-poison-1826)

[ecto basics](https://elixirschool.com/pt/lessons/ecto/basics/)