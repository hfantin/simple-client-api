#!/bin/bash

function main(){
    echo "> Executa build em todos os projetos"
    for dir in */; do
      echo "> $dir"
      if [[ -e "$dir/Dockerfile" ]]; then
        if [[ -e "$dir/build.gradle" ]]; then
            echo 'building springboot app'
            cd $dir
            ./gradlew clean build docker
            cd ..
        fi
      fi
    done
}

main
