FROM elixir:1.9.2-alpine
WORKDIR /opt/app
COPY . .
RUN mix local.hex --force
RUN mix deps.get
CMD ["mix", "phx.server"]
