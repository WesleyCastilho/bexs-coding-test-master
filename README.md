# Bexs Bank - Software Engineer Rota de Viagem
Rota de Viagem

Um turista deseja viajar pelo mundo pagando o menor preço possível independentemente do número de conexões necessárias. Vamos construir um programa que facilite ao nosso turista, escolher a melhor rota para sua viagem.

Para isso precisamos inserir as rotas através de um arquivo de entrada.

## Como utilizar

- Confirme a versão instalada da linguagem Go, para este projeto utilizamos a versão 1.16

```bash
go version
```

Bibliotecas serão instaladas automaticamente ao executar o projeto

### Execução por linha de comando
- Para execução via terminal (CLI), execute o comando abaixo, informando o caminho para o arquivo .csv contendo as informaçõea de rotas de viagem:

```bash
go run ./cli/input-routes.csv
```

Após isto, basta seguir as instruções do prompt, mas lembre-se a rota digitada para consulta degve seguir o padrão IATA para aeroportos, por exemplo: CGH-FLN (Congonhas-Florianópolis)
```bash

Abaixo mostramos como podem ser realizadas as consultas.
=========================================================================================
Iniciando.... Banco Bexs Melhor Rota, para sair digite 'Sair' ou, pressione ctrl+C.
=========================================================================================

Por favor digite a rota, no padrão ORI-DES: : GRU-CDG
melhor preço: GRU - BRC > $10
Por favor digite a rota, no padrão ORI-DES: : GRU-SCL
melhor preço: GRU - BRC - SCL > $15
Por favor digite a rota, no padrão ORI-DES: : GRU-ORL
melhor preço: GRU - BRC - ORL > $16
Por favor digite a rota, no padrão ORI-DES: : SCL-ORL
melhor preço: SCL - ORL > $20
Por favor digite a rota, no padrão ORI-DES: : BRC-ORL
melhor preço: BRC - ORL > $6

Para os casos de erro, favor atentar-se as instruções
Por favor digite a rota, no padrão ORI-DES: : BRC-GRU
Rota não encontrada, BRC->GRU. por favor, tente novamnte
Por favor digite a rota, no padrão ORI-DES: : BRCC
entrada inválida BRCC. o formato deve ser XXX-XXX, tente novamente
Por favor digite a rota, no padrão ORI-DES: : BRCC-GRUU
origem BRCC não encontrada. tente informar uma origem existente ou crie uma rota com essa origem.
Por favor digite a rota, no padrão ORI-DES: : GRUU-BRCC
origem GRUU não encontrada. tente informar uma origem existente ou crie uma rota com essa origem.
Por favor digite a rota, no padrão ORI-DES: : 
entrada inválida . o formato deve ser XXX-XXX, tente novamente
Por favor digite a rota, no padrão ORI-DES: : _-_
origem não encontrada. tente informar uma existent origem ou crie uma rota com essa origem.
Por fa vor digite a rota, no padrão ORI-DES: : sair

```

### API REST

-  Para executar o servidor, digite comando abaixo, o servidor http vai subir na porta 8080 e fará a leitura do arquivo csv em:  data/input-routes.csv

```bash
go run ./rest
```


## O servidor REST expõe 3 rotas:

- GET /: A rota '/' é a rota de entrada na API. Foi pensada para informar o status geral da API. <br/>
Retorna um header com status 200, caso o servidor esteja on line

Exemplo:

```bash
curl -i -GET "localhost:8080"
```

- GET /route: Procura a melhor rota baseado nos parâmetros recebidos pela url (query params) (ex. 'localhost:8080/route?from=GRU&to=BRC' irá procurar a melhor rota (cadastrada) entre Guarulhos/BR e San Carlos de Bariloche/AR:
  - from (obrigatório - required): código de origem (GRU, BRC.. etc sempre no formato [IATA](http://atualdespachos.com.br/assets/Uploads/Aeroportos-do-Mundo-e-cdigo-IATA.pdf) )
  - to (obrigatório - required): código de destino, possui as mesmas regras do código de origem
  - Retorna um objeto JSON contendo a origem, destino, a rota inteira (com as conexoes entre origem e destino) e o preço total da rota.<br/>
    {"From":"GRU","To":"CDG","Path":"GRU - CDG","Price":10}). <br/>
    * Se houver um erro, retorna com o status http error 

Example:

```bash
curl -i -GET "localhost:8080/route?from=GRU&to=BRC"
```

- POST /route: Cadastra uma nova rota, passando os valores dos argumentos em um objeto JSON:
  - from (obrigatório - required, string): código de origem (ex. GRU, CDG)
  - to (obrigatório - required, string): código de destino (ex. GRU, CDG)
  - price (obrigatório - required, int): path price (e.g. 20, 5)
  - Retorna um header com status 200 se for processado com sucesso. Se houver um erro, retorna com o status http error 

Example:

```bash
curl -i -XPOST "localhost:8080/route" -H "Content-Type: application/json" --data '{"from": "BLL", "to": "SXF", "price": 10}'
```



### Testes

A definição do projeto determina que seja possível utilizar via linha de comando e via API REST, dessa forma separei a lógica de negócios em 2 modulos, cli e rest. 

Todo o projeto é testável, e pode ser feito de forma geral, ou por partes

Na raiz do projeto, execute: 

Para testar todo o código de uma vez
```bash
go test ./... -v
$HOME/go/bin/gotest ./... -v
```

Para testar strings e parsers:

```bash
go test ./services/utils -v
$HOME/go/bin/gotest ./services/utils -v
```

Para testar os arquivos csv:

```bash
go test ./services/csv -v
$HOME/go/bin/gotest ./services/csv -v
```

Para o teste de dominio:

```bash
go test ./domain -v
$HOME/go/bin/gotest ./domain -v
```


## Estrutura de arquivos e pastas
    .
    ├─ cli                     # Interface de linha de comando, para usar no terminal
    ├─ coverage                # Arquivo dos testes de cobertura, gerados pelo Go.
    ├─ data                    # Arquivo contendo csv para consulta de rota, inicialmente possui o que veio na definição do projeto
    ├─ domain                  # Dominio da Aplicação
    ├─ rest                    # Interface e servidor HTTP, para utilizacao via REST
    ├─ services                # Pacote com ferramentas para auxilio (utils) na tratativa com leitura e gravação de arquivos, sendo necessario para criação de novas rotas.
    ├─ go.mod                  # Arquivo de dependências do projeto.
    └─ README.md               # Descrição e informações sobre o projeto

## Decisões sobre a Resolução do Problema

### Brief algorithm explanation

O teste se assemelha com o "Problema do Caixeiro Viajante", que visa encontar a menor distancia entre um ponto e outro(cidades no caso) sem repetir o caminho(rota)

Para solucionar foi utilizado o algoritmo de Dijkstra (https://pt.wikipedia.org/wiki/Algoritmo_de_Dijkstra), este algoritmo recebe um grafo
orientado (G,w) (sem arestas de peso negativo) e um vértice s de G, e devolve para cada v E V[G], o peso de um caminho mínimo de s a v.

O algoritmo utiliza recursão, sempre buscando o caminho mais barato procurando dentro dos nós internos e identificando se é um destino possível ou se o caminho interno deve ser descartado (semelhante à primeira pesquisa de profundidade nas pesquisas em árvore).
Ao longo da execução, os caminhos são acumulados em uma matriz, enquanto o preço total é aumentado de acordo com o preço da rota mais barata.
Para aplicar o principio de responsabilidade unica, a logica da rota principal, foi separada em um pacote e exposta como uma interface para utilizacao tanto pelo cli quanto pela API.
Para dividir as responsabilidades, a lógica da rota principal foi separada em um pacote, exposta como uma interface para a interface de linha de comando e o servidor HTTP REST. 
Da mesma forma, outras responsabilidades foram agrupadas em pacotes de serviço, como um analisador CSV para ler/gravar o arquivo de rotas de entrada, analisadores de string/arquivo que combinam métodos em utilitários (por exemplo, trim e superior uma string para a origem / comparação do código de destino).

### Decisions over Go

A opção pela linguagem Go, foi além de passional(rs) pautadas nos seguintes pontos:

- Quase baixo nível, eu aprendi programar em C, então é fácil perceber que "objetos" em Go são mais leves devido à falta de alta hierarquia, com indicações diretas, como ponteiros

- A limitação com relação a utilização de frameworks(no desafio), com Go conseguimos fazer quase tudo de forma nativa, deixando o código limpo. 

- Go é moderno, foi criado para aproveitarmos o melhor dos equipamentos atuais, como os processadores multi-core, evitando o uso de memória compartilhada e dispensando o uso de travas, semáforos e outras técnicas de sincronização de processos. 
Além de possuir uma extensa biblioteca padrão com ferramentas para comunicação em redes, servidores HTTP, expressões regulares, leitura e escrita de arquivos.

- Compila pra qualquer coisa (praticamente)

- Flexibilidade: Não preciso me locomover entre múltiplas interfaces, estruturas super-hierárquicas ou código verboso, Go trabalha bem com múltiplos  paradigmas por meio de suas composições, estruturas e tratamento de erros

- Go é orientado para micro serviços, como estou procurando alinhar a minha carreira com o modelo Full Cycle, e aplicações em arquitetura hexagonal,
 fez mais sentido para mim, além de ser parte da Stack utilizada atualmente pelo Banco Bexs [:)] 
 
### Servidor HTTP

Para o servidor HTTP, foi utilizado o pacote net/http.
As rotas foram divididas em duas funções diferentes (ApiStatus e ServeHTTP), onde o ServeHTTP que serve o caminho / rota tem uma composição de wrapper em torno da estrutura * domain.Routes *. 
Isso permite uma manipulação diretamente para a classe de estrutura de rotas do pacote de domínio, permitindo que o servidor faça pesquisas de maneira adequada e insira novas rotas no arquivo * input-routes.csv *.

### CLI

Para a interface da linha de comandos, não houve nada de especial em sua implementação. 

Verificamos o formato dos argumentos, e a presença do input de rotas, o terminanl continua ativo, até que o usuário digite ** exit ** ou tecle ** Ctrl + C **, comandos que vão encerrar o processo.

### Melhorias

Uma aplicação nunca está pronta, embora funcional, há diversos pontos que podem/devem ser melhorados, que não foi feito por falta de tempo ou conhecimento.

- Iterligar os testes para o cli e para o server http

- Trabalhar melhor as rotas (verbos) não implementados (501,404, por exemplo)

- Refatoração do código da camada de domínio, para aumentar a coesão do código

- Bloquear o acesso aos arquivos csv, enquanto estiverem sendo lidos

- Persistir as rotas em um banco de dados (sqlite porexemplo)