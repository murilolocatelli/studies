(ns curso-clojure.aula4)

(println "Aula 4")

(def precos [10 20 30 150 160])

(println (get precos 1))
(println (precos 0))

(println "valor padrão nil" (get precos 5))
(println "valor padrão 0" (get precos 5 0))
; (println (precos 5)) IndexOutOfBoundsException

(println "cria um novo vetor e retorna (imutabilidade)" (conj precos 5))
(println "vetor original" precos)

(println (+ 5 1))
(println (inc 5))
(println (update precos 0 inc))
(println (update precos 0 dec))

(defn soma-1
  [valor]
  (+ valor 1))

(println (update precos 0 soma-1))

(defn aplica-desconto?
  [valor-bruto]
  (> valor-bruto 100))

(defn valor-descontado
  "Retorna o valor com desconto"
  [valor-bruto]
  (if (aplica-desconto? valor-bruto)
    (* valor-bruto 0.9)
    valor-bruto))

(println "map:" (map valor-descontado precos))

(println "filter:" (filter even? (range 10)))
(println "filter:" (filter aplica-desconto? precos))
(println "map / filter:" (map valor-descontado (filter aplica-desconto? precos)))

(defn minha-soma
  "Minha Soma"
  [valor1, valor2]
  (println "somando" valor1 valor2)
  (+ valor1 valor2))

(println "reduce:" (reduce + precos))
(println "reduce com valor default:" (reduce minha-soma 0 precos))
