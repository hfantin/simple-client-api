# clojure-rest-api

Projeto basico criado para testes de performance entre linguagens com requisição rest e acesso a base de dados MySQL.
Feito em conjunto com [@hfantin]: https://github.com/hfantin/simple-client-api

## Prerequisites

You will need [Leiningen][] 2.0.0 or above installed.

[leiningen]: https://github.com/technomancy/leiningen

## Running

To start a web server for the application, run:

    lein ring server

If you need to start with another port use env variables:
    
    PORT=5000 lein ring server
