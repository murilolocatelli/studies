(ns curso-clojure.aula6)

(def pedido {:mochila {:quantidade 10, :preco 20}
             :cadeira {:quantidade 10, :preco 30}})

(defn imprime
  [param]
  (println "param class:" (class param))
  (println "param count:" (count param))
  (println "param key:" (get param 0))
  (println "param value:" (get param 1)))

(println "1:" (map imprime pedido))

(defn imprime
  [[key value]]
  (println "key:" key)
  (println "value:" value))

(println "2:" (map imprime pedido))

(defn preco-dos-produtos
  [[_ value]]
  (* (:quantidade value) (:preco value)))

(println (reduce + (map preco-dos-produtos pedido)))

; threading last
(defn total-do-pedido
  [pedido]
  (->> pedido
       (map preco-dos-produtos)
       (reduce +)))

(println "threading last:" (total-do-pedido pedido))

; passando somente o value do mapa
(defn preco-do-produto
  [produto]
  (println "produto:" produto)
  (* (:quantidade produto) (:preco produto)))

(println (reduce + (map preco-do-produto (vals pedido))))

; threading last
(defn total-do-pedido
  [pedido]
  (->> pedido
       vals
       (map preco-do-produto)
       (reduce +)))

(println "threading last 2:" (total-do-pedido pedido))

(def pedido {:mochila {:quantidade 10, :preco 20}
             :cadeira {:quantidade 10, :preco 30}
             :chaveiro {:quantidade 1}})

(defn gratuito?
  [item]
  (<= (get item :preco 0) 0))

(println "gratuitos:" (filter gratuito? (vals pedido)))
(println "gratuitos 2:" (filter (fn [[chave valor]] (gratuito? valor)) pedido))
(println "gratuitos 3:" (filter #(gratuito? (second %)) pedido))

(defn pago?
  [item]
  (not (gratuito? item)))

(println "pago:" (pago? {:preco 1}))
(println "pago:" ((comp not gratuito?) {:preco 1}))

(def pago (comp not gratuito?))
(println "pago:" (pago {:preco 1}))
