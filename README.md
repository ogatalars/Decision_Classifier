# Decision_Classifier
Projeto pensado para classificar as ementas de decisões extraídas das páginas dos tribunais de justiça de todo país em categorias:

- Prejudicado
- Convertido em diligência
- Reexame necessári
- Parcialmente Procedente
- Improcedente
- Procedente
- Sem Ementa
- Casos não mapeados

É possível fazer uso apenas da função isolada de classificação que retorna apenas uma sequência em string, exemplo:

```
summary := " CONDOMÍNIO. COBRANÇA. [...]. RECURSO PROVIDO. "
id := "0"
court := "nowhere"

inferedDecision := Classifier.Decision_Classifier(summary, id, court)

```

Como também é possível realizar a leitura de um csv contendo as ementas de decisões e serão retornados 9 arquivos csv apartados correspondetes a uma classificação mencionada acima, sendo esses:

- affected.csv
- diligence.csv
- exOfficioReview.csv
- partial.csv 
- groundless.csv 
- hasGrounds.csv 
- noSummary.csv
- notMapped.csv 
- Total_Classification <> **Compilado de todos os documentos acima**

O csv com os dados deve ter a sequência {id, identifier, decision, court} e para utilizar o leitor e classificador basta utilizar a função Loop:

```
CSV.Loop(CSV.ReadCsvFile("CSV/Database/testData.csv"))
```

## Run
```go run main.go```

## Disclaimer
Esse é um work in progress do classificador, ainda não sabemos o % de efetividade desse classificador para as ementas existentes nos tribunais de justiça. Use por sua conta em risco.

## Considerações
A) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.

B) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
