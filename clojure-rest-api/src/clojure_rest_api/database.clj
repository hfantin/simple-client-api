(ns clojure-rest-api.database
    (:require [korma.db :as korma]))
  
  (def db-connection-info 
      (korma/mysql 
          {:classname "com.mysql.cj.jdbc.Driver"
          :subprotocol "mysql"
          :user "guest"
          :password "guest"
          :subname "//localhost:3306/test"}))
  
  ; set up korma
  (korma/defdb db db-connection-info)