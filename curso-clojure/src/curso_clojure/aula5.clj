(ns curso-clojure.aula5)

(def estoque {"mochila" 10 "camiseta" 5})
(println estoque)

(def estoque {"mochila" 10 "camiseta" 5})
(println "count:" (count estoque))

(def estoque {"mochila" 10
              "camiseta" 5})
(println "keys:" (keys estoque))
(println "values:" (vals estoque))

(def estoque {:mochila 10
              :camiseta 5})
(println "associa (novo mapa):" (assoc estoque :cadeira 3))
(println "sobrescreve chave (novo mapa):" (assoc estoque :mochila 1))

(println "update (novo mapa):" (update estoque :mochila inc))
(println estoque)

(println "desassocia (novo mapa):" (dissoc estoque :mochila))

(println "pedido:")
(def pedido {:mochila {:quantidade 10, :preco 20}
             :cadeira {:quantidade 10, :preco 30}})
(def pedidoNil nil)
(println pedido)
(println (pedido :mochila))
(println (pedido :sapato))
(println (get pedido :mochila))
(println (get pedido :sapato))
(println (get pedido :sapato {}))
(println (:mochila pedido))
(println (:sapato pedido))
(println (:sapato pedido {}))
; (println (pedidoNil :sapato)) NullPointerException
(println (get pedidoNil :sapato))
(println (:mochila pedidoNil))

(println "Forma funcional de navegar" (:quantidade (:mochila pedido)))
; Threading first
(println "Forma tradicional de navegar" (-> pedido
                                            :mochila
                                            :quantidade))
(-> pedido
    :mochila
    :quantidade
    println)

(println (update-in pedido [:mochila :quantidade] inc))


