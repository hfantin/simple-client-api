(defproject clojure-rest-api "0.1.0-SNAPSHOT"
  :description "Projeto base de requisições rest para testes de performance multilinguagem"
  :url "https://github.com/alelfreitas"
  :min-lein-version "2.0.0"
  :dependencies [[org.clojure/clojure "1.10.0"]
                 [compojure "1.6.1"]
                 [ring/ring-core "1.7.1"]
                 [ring/ring-json "0.5.0"]
                 [korma "0.4.3"]
                 [mysql/mysql-connector-java "8.0.17"]
                 [ring/ring-defaults "0.3.2"]]
  :plugins [[lein-ring "0.12.5"]]
  :ring {:handler clojure-rest-api.handler/app}
  :profiles
  {:dev {:dependencies [[javax.servlet/servlet-api "2.5"]
                        [ring/ring-mock "0.3.2"]]}})
