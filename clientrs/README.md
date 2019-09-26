# simple client api
- Rust with Rocket and Diesel

based on [this article](https://cprimozic.net/blog/rust-rocket-cloud-run/)
https://lankydan.dev/2018/05/20/creating-a-rusty-rocket-fuelled-with-diesel


# Instalation
rust and libs:   
- sudo apt-get install gcc libmysqlclient-dev libpq-dev
- curl https://sh.rustup.rs -sSf | sh
- rustup install nightly
- rustup default nightly

diese cli:    
- cargo install diesel_cli --no-default-features --features postgres mysql 

environment variables:   
- mysql:   export DATABASE_URL=mysql://guest:guest@localhost:3306/test
           echo DATABASE_URL=mysql://guest:guest@localhost:3306/test > .env
- postgres: export DATABASE_URL=postgres://guest:guest@localhost/test 


diesel  
- diesel setup # creates migration directory
- diesel migration generate initialize
- diesel migration run

# Run
export ROCKET_DATABASES="{ test = { url = \"$DATABASE_URL\" } }"
cargo run