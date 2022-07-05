### Reclutamiento de mutantes

####Prerequisito
- Descargar el repositorio del desarrollo a un directorio local


####Testing
Para la ejecucion de las pruebas del desarrollo realizar los siguientes pasos:
- Ejecutar desde la terminal el siguiente comando:
**go test ./... -coverprofile=coverage.out**
Mostrara el resultado de los test automaticos y code coverage

- Ejecutar desde la terminal el siguiente comando:
**go tool cover -html=coverage.out**
Abrira una pagina html con la cobertura del test automatico

####Docker
Ubicarse en la carpeta **database** del repositorio y ejecutar los siguientes comandos:
**docker build . -t recruitent_mutant-ws-rest-db** 
**docker run -p 54321:5432 recruitent_mutant-ws-rest-db**

####Revision del API REST
Donde se pueda detectar si un humano es mutante enviando la secuencia de ADN mediante un HTTP POST con un Json el cual tenga el siguiente formato:
> {
  "dna": [
    "ATGCTGCTAC",
    "TGTTCTGTTT",
    "GAAGCCCTAT",
    "CGGGTTGCCG",
    "ATGCTGGAAC",
    "GGTTAAGGCG",
    "TTGGAACTCG",
    "CCGCGGCCTC",
    "GGTGTCTGCG",
    "GTTTTCTTGC"
  ]
}
>