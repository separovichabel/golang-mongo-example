# Desafio de dados

Esse desafio de dados consiste em provisionar dados para 3 finalidades diferentes. Suas especificações são:

* Base A:
    * Dados sensiveis
    * Sem necessidade de acesso muito performático
* Base B:
    * Dados sensiveis
    * Necessidade de acesso mais performático que a base A
    * Multiplos usuarios do sistema
    * Tambem usado para Machine Learning
* Base C:
    * Dados não sensiveis
    * Necessidade de acesso muito performático

## Explicações gerais 
Dado o tempo para a solução do problema, irei explicar possiveis soluções para as 3 bases de dados e implementar apenas 1.

## Questões gerais de segurança
Para resolver os problemas de segurança paços padrões de segurança podem ser seguidos.

### A nivel de banco:

* Ter bom controle do uso das senhas do banco (usar senhas fortes e guardadas adequadamente).
* Ter regras de acesso, criação, alteração e deleção de dados bem definidos por usuario e aplicação.
* Usar um Proxy / Firewall para ter controle de cada query executada.
* Usar criptografia de ponta a ponta para ter conexão segura entre aplicação e base (aplicação -> proxy e proxy -> banco).
* Ter backups de rotina e seguros.
* Usar redes privadas para o uso do banco para não exposto-lo a internet e possiveis acessos externos.
* CASO NECESSARIO, criptografar ou alterar dados sensiveis.

### A nivel de Aplicação:

* Usar HTTPS para criptografia no trafego dos dados.
* Usar ferramentas robustas de autenticação e autorização (JWT e OAuth2 por exemplo).
* Exigir mais de 1 fator de autenticação. 
* Delimitar bem o grupos de acessos.
* Usar bons padrões de codigo para acesso de dados (query builders que impedem SQL injection ou ORMs).
* Para um nivel mais alto de segurança, ter renovação e expiração de tokens com curto espaço de tempo (no caso de comunicação usuario -> aplicação)
* Nunca guardar senhas no codigo, mas usar cofres e injetar senhas em tempo de execução (nos containers).
* Ter ambientes separados de teste e produção, com credenciais e dados diferentes.

### A nivel de Infraestrutura:

* Criar redes privadas para trafego interno de dados.
* Usar Firewall e outras ferramentas para controle de acesso externo a rede privada.
* Ter segregação de camadas de aplicação e banco em maquinas diferentes. 

## Sugestões para a Base A
Modelagem sugerida (exemplo): 

![Modelagem Base A](images/Modelagem%20Base%20A.png)

Dado que nos requisitos da Base A não incluem alta demanda e nem performance, uma base de dado simples com foco em segurança creio ser o suficiente. Algo parecido com o desenho a seguir:

![Arquitetura sugerida Base A](images/Arquitetura%20sugerida%20base%20A.png)

## Sugestões para a Base B
Modelagem sugerida (exemplo): 

![Modelagem Base A](images/Modelagem%20Base%20B.png)

Dado que os mesmos passos de segurança podem ser aplicados tambem para a Base B, os mesmos passos não serão abordados aqui, e serão ignorados (pré supostos) nos desenhos de arquitetura.

Visto que a demanda na Base B exige performance com multiplos usuarios, adotar uma estratégia de clusterização com replicas dedicadas a leitura seria uma boa opção para melhorar o tempo de leitura.

Além disso, para a disponibilização dos dados para machine learning, pode ser avaliado a remoção ou criação de dados FAKE no caso de dados sensiveis.
Por exemplo, ao buscar os dados de uma pessoa, o valor do patrimonio é um dado relevante, mas o nome da pessoa é irrelevante. A remoção do nome da pessoa pode ser feita, ou a sua troca por um nome fake.

Para essa finalidade poderia ser criada uma base replica da Base B, porém sem dados sensiveis guardados. E todo dado inserido na Base B seria convertido para a Replica Segura Base B.

Uma possivel implementação seria:

![Arquitetura sugerida Base A](images/Arquitetura%20sugerida%20base%20B.png)

Assim o acesso aos dados sensiveis continuam com as mesmas camadas de segurança. E para uma outra ferramenta sem dados sensiveis uma eventual ferramenta de ML poderia acessar a base diretamente (possivelmente).


## Sugestões para a Base C
A Base C não precisa de muitas ferramentas de segurança, porém seus requisitos de velocidade são primordiais.

Um ponto positivo a ser levado em consideração é que os dados da Base C não são históricos, mas apenas é armazenado o ultimo evento de cada tópico, que são:

- Última Consulta do CPF Bureau
- Movimentacao financeira ( última )
- Última compra cartao

Como os dados aparentemente só se referem a ultimas atividades, vou adotar tambem a movimentação financeira também se referindo a última.

Portanto, não havendo necessidade histórica, e necessaria a alta disponibilidade, e a pesquisa ser esclusivamente relacionada a um CPF, uma implementação de uma banco Chave - Valor parece ser a melhor opção.

A pesquisa é sempre feita apenas por CPF ( chave ), e os dados retornados são sempre todos (trazendo sempre todos os dados, sem a possibilidade de filtrar apenas movimentação financeira).

Isso permite a sugestão da seguinte modelagem:

![Modelagem Base A](images/Modelagem%20Base%20C.png)

Para essa solução pode ser usado um banco como MongoDB ou DynamoDB, ou até mesmo um banco de Cache como banco de dados primário.

![Arquitetura sugerida Base A](images/Arquitetura%20sugerida%20base%20C.png)

Esse será o problema que eu irei resolver, por isso terá uma documentação dedicada.