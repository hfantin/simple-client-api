FROM elixir:latest
WORKDIR /opt/app
COPY . .
RUN mix local.hex --force
RUN mix deps.get
CMD ["mix", "phx.server"]
