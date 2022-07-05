### Reclutamiento de mutantes

#### Prerequisito
- Descargar el repositorio del desarrollo a un directorio local


#### Testing
Para la ejecucion de las pruebas del desarrollo realizar los siguientes pasos:
- Ejecutar desde la terminal el siguiente comando:
**go test ./... -coverprofile=coverage.out**
Mostrara el resultado de los test automaticos y code coverage

- Ejecutar desde la terminal el siguiente comando:
**go tool cover -html=coverage.out**
Abrira una pagina html con la cobertura del test automatico

#### Docker
Ubicarse en la carpeta **database** del repositorio y ejecutar los siguientes comandos:
**docker build . -t recruitent_mutant-ws-rest-db** 
**docker run -p 54321:5432 recruitent_mutant-ws-rest-db**

#### Revision del API REST
Donde se pueda detectar si un humano es mutante enviando la secuencia de ADN mediante un HTTP POST con un Json el cual tenga el siguiente formato:

<div class="highlight highlight-source-json notranslate position-relative overflow-auto"><pre>{
  <span class="pl-ent">"dna"</span>: [
    <span class="pl-s"><span class="pl-pds">"</span>ATGCGA<span class="pl-pds">"</span></span>,
    <span class="pl-s"><span class="pl-pds">"</span>CAGTGC<span class="pl-pds">"</span></span>,
    <span class="pl-s"><span class="pl-pds">"</span>TTATGT<span class="pl-pds">"</span></span>,
    <span class="pl-s"><span class="pl-pds">"</span>AGAAGG<span class="pl-pds">"</span></span>,
    <span class="pl-s"><span class="pl-pds">"</span>CCCCTA<span class="pl-pds">"</span></span>,
    <span class="pl-s"><span class="pl-pds">"</span>TCACTG<span class="pl-pds">"</span></span>
  ]
}</pre>
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-check js-clipboard-check-icon color-fg-success d-none m-2" wfd-invisible="true">
    <path fill-rule="evenodd" d="M13.78 4.22a.75.75 0 010 1.06l-7.25 7.25a.75.75 0 01-1.06 0L2.22 9.28a.75.75 0 011.06-1.06L6 10.94l6.72-6.72a.75.75 0 011.06 0z"></path>
</svg>
    </clipboard-copy>
  </div>