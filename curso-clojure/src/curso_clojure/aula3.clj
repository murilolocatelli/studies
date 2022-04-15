(ns curso-clojure.aula3)

; Define uma função
(defn aplica-desconto?
  [valor-bruto]
  (if (> valor-bruto 100)
    true))

; Define um símbolo que recebe uma função anônima
(def aplica-desconto-2? (fn [valor-bruto] (> valor-bruto 100)))

(println (aplica-desconto? 100))
(println (aplica-desconto-2? 1000))

(defn valor-descontado
  "Retorna o valor com desconto"
  [aplica? valor-bruto]
  (if (aplica-desconto? valor-bruto)
    (* valor-bruto 0.9)
    valor-bruto))

(println "Função como parâmetro")
(println (valor-descontado aplica-desconto? 100))

(println "Função anônima")
(println (valor-descontado (fn [valor-bruto] (> valor-bruto 100)) 1000))
(println (valor-descontado #(> %1 100) 1000))
(println (valor-descontado #(> % 100) 1000))
