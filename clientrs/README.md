# simple client api
- Rust with Rocket and Diesel

based on [this article](https://cprimozic.net/blog/rust-rocket-cloud-run/)

# Instalation
- sudo apt-get install gcc
- curl https://sh.rustup.rs -sSf | sh
- rustup install nightly
- rustup default nightly
- cargo install diesel_cli --no-default-features --features mysql


# Run
export ROCKET_DATABASES="{ clientrs = { url = \"guest:guest@localhost:3306/test\" } }"
cargo run

