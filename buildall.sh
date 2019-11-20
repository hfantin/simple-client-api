#!/bin/bash

function main(){
    echo "> Building all projects..."
    for dir in */; do
      echo "> $dir" 
      app="github.com/hfantin/${dir%/}"
      if [[ -e "$dir/Dockerfile" ]]; then
        echo "building $app"
        cd $dir
        if [[ -e "main.go" ]]; then
          GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/app 
          docker build -t "$app" .
        # elif [[ -e "build.gradle" ]]; then
        #     ./gradlew clean build docker
        # elif [[ -e "package.json" ]]; then
        #    npm i
        #    docker build -t "$app" .
        # elif [[ -e "mix.exs" ]]; then
        #    docker build -t "$app" .
        # elif [[ -e "Cargo.toml" ]]; then
        #    cargo build --release
        #    docker build -t "$app" .
        fi
        cd ..
      fi
    done
}


main

# TODO clientrs with alpine
#           rustup target add x86_64-unknown-linux-musl
#         #  crosscompile to run in alpine
#         #  apt install musl-gcc
#         #   cargo build --target x86_64-unknown-linux-musl --release
#           # musl-gcc 
#           cargo build --release --target=x86_64-unknown-linux-musl
#           # cargo build --release --target=x86_64-unknown-linux-musl