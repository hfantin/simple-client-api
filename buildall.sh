#!/bin/bash

function main(){
    echo "> Executa build em todos os projetos"
    for dir in */; do
      echo "> $dir"
      if [[ -e "$dir/Dockerfile" ]]; then
        echo "aqui tem dockerfile $dir"
        if [[ -e "$dir/build.gradle" ]]; then
            cd $dir
            echo 'building springboot app'
            # ./gradlew clean build docker
            cd ..
        fi
        if [[ -e "$dir/main.go" ]]; then
          cd $dir
          echo 'building golang app'
          # go build -o main .
          go build
          docker build -t "clientgov4" .
          cd ..
        fi
      fi
    done
}

main
