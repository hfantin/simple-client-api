# simple client api
- Rust with Rocket and Diesel

based on [this article](https://cprimozic.net/blog/rust-rocket-cloud-run/)
https://lankydan.dev/2018/05/20/creating-a-rusty-rocket-fuelled-with-diesel


# Instalation
- sudo apt-get install gcc libmysqlclient-dev libpq-dev
- curl https://sh.rustup.rs -sSf | sh
- rustup install nightly
- rustup default nightly
- cargo install diesel_cli --no-default-features --features mysql
- export DATABASE_URL=mysql://guest:guest@localhost:3306/test
  export DATABASE_URL=postgres://guest:guest@localhost/test 
- diesel setup
- diesel migration generate initialize
- diesel migration run


# Run
export ROCKET_DATABASES="{ clientrs = { url = \"guest:guest@localhost:3306\" } }"
export DATABASE_URL="mysql://guest:guest@tcp(localhost:3306)/test"
echo DATABASE_URL=mysql://guest:guest@localhost:3306/test > .env
cargo run






-- Drop table

-- DROP TABLE test.clients;

CREATE TABLE test.clients (
	id serial NOT NULL,
	"name" varchar NOT NULL,
	email varchar(100) NULL
);
