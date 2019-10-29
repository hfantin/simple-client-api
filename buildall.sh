#!/bin/bash

function main(){
    echo "> Executa build em todos os projetos"
    for dir in */; do
      echo "> $dir"
      if [[ -e "$dir/Dockerfile" ]]; then
        echo "aqui tem dockerfile $dir"
        cd $dir
        if [[ -e "build.gradle" ]]; then
            echo 'building springboot app'
            # ./gradlew clean build docker
        elif [[ -e "main.go" ]]; then
          echo 'building golang app'
          GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o clientgo
          # docker build -t "clientgo" .
        elif [[ -e "package.json" ]]; then
           echo 'building nodejs app'
          #  docker build -t "clientjs" .
        elif [[ -e "mix.exs" ]]; then
           echo 'building elixir app'
           docker build -t "clientex" .
        elif [[ -e "Cargo.toml" ]]; then
           echo 'building rust app'
           cargo build --release
           docker build -t "clientrs" .
#           rustup target add x86_64-unknown-linux-musl
#         #  crosscompile to run in alpine
#         #  apt install musl-gcc
#         #   cargo build --target x86_64-unknown-linux-musl --release
#           # musl-gcc 
#           cargo build --release --target=x86_64-unknown-linux-musl
#           # cargo build --release --target=x86_64-unknown-linux-musl
        fi
        cd ..
      fi
    done
}

main