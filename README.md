## Introdução
1. Para realizar pesquisas de NPS (https://satisfacaodeclientes.com/net-promoter-score/) é preciso levar em conta algumas regras, como por exemplo, a Regra da Noventena. Esta regra diz que um cliente não deve ser impactado em menos de 90 dias mais de uma vez pela pesquisa de NPS. Com essa regra é possível garantir que o cliente não será incomodado com muitas perguntas, o que geraria atrito entre o cliente e a empresa. Isso é chamado de Fadiga de Pesquisa, algo comum nas empresas que fazem pesquisas com tecnologias precárias.   
 A Tracksale, sendo referência em pesquisas de NPS, se preocupa muito com a relação que nossos clientes tem com os seus consumidores, por isso, a Regra de Noventena é um dos pontos críticos da nossa aplicação. Sabendo que para checar se um cliente deve ser impactado por uma pesquisa é um processo muito custoso, precisamos desenvolver uma forma segura de contornar essa checagem. 

## Requisitos de tecnologia

* Desenvolver um projeto que tenha um endpoint onde seja possível verificar se um cliente deve ou não ser impactado por uma pesquisa.
Toda vez que a checagem retornar positivo (o usuário deve ser impactado pela pesquisa), o cliente que será impactado deverá ser incluído na Regra da Noventena.
* O resultado da checagem deverá ser mantida em um cache que deverá ser invalidado quando necessário.
* A aplicação desenvolvida deverá estar coberta por testes.
* Realize o teste na linguagem de sua preferência, porém, contará pontos se desenvolvida em alguma das linguagens abaixo:
 - Golang
 - Elixir
 - C/C++

## Sobre o Projeto
* Foi utilizado as seguintes tecnologias:
    - Docker como container
    - Golang e Redis.

* Golang
    - Implementado a rotina de controle do NPS.

* Redis
    - Servidor responsável por armazenar a última data de envio de NPS de cada cliente.
    - https://docker-doc.readthedocs.io/zh_CN/stable/examples/running_redis_service.html

## Artefatos entregues
* Validação de NPS
* Tests

## Subindo a aplicação
* build
    - Fazer o checkout do repositório e entrar no mesmo.
    - Rodar o comando 'docker-compose up -d'

* Rodando os tests da aplicação
    - Com o container rodando, entrar no container com o comando 'docker exec -ti [CONTAINER] /bin/sh'
    - Após entrar no container rodar o comando 'go test'

## Acessando a aplicação
* Com o container rodando acessar o endereço 'http://localhost:6060/?key={IDCLIENTE}'
    - IDCLIENTE é o cliente que deseja verificar se tem que enviar NPS

## Retorno da Requição
* Liberado para o envio de avaliação!
* Não está liberado para o envio de avaliação!