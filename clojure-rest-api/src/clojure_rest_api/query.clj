(ns clojure-rest-api.query
    (:require [clojure_rest_api.database]
              [korma.core :refer :all]))
  
  (defentity CLI)
  
  (defn get-clients []
    (select CLI))
  
  