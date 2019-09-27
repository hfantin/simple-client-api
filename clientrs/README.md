# Creating database tables

1. curl https://sh.rustup.rs -sSf | sh
2. rustup install nightly
3. rustup default nightly
4. sudo apt-get install gcc libmysqlclient-dev libpq-dev
5. cargo install diesel_cli --no-default-features --features postgres
6. echo DATABASE_URL=postgres://guest:guest@localhost/test > .env
7. diesel setup
8. diesel migration generate create_people
9. diesel migration run -> to apply migration
   diesel migration redo -> to undo migration
10. diesel print-schema > src/schema.rs -> to generate schema, if there is no one

# links

https://github.com/lankydan/rust-web-with-rocket
https://lankydan.dev/2018/05/20/creating-a-rusty-rocket-fuelled-with-diesel
