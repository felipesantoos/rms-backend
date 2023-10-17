# GO Backend Template

## O que é o projeto?

Esse é um projeto que utiliza a arquitetura hexagonal e até o momento tem uma única porta de entrada: um projeto BackEnd que utiliza o Framework [Echo](https://echo.labstack.com/).

## Como executar?

Os passos abaixo assumem que você já fez/tem:
1. O clone (`git clone`) do projeto no seu computador e tem um terminal aberto na pasta baixada;

Caso você seja desenvolvedor:

2. A ferramenta para CLI do `go` instalada (caso não tenha, clique [aqui](https://go.dev/learn/) para ir para a documentação de instalação)

Caso você seja do time de QA:

3. A ferramenta docker instalada (caso não tenha, clique [aqui](https://www.docker.com/) para ir para a página de instalação)

### **Para Desenvolvimento**

Se você é um membro do time de desenvolvimento desse projeto, siga os passos abaixo para executar as configurações apropriadas:

1. Execute o comando `go mod tidy` para baixar as dependências do projeto;
2. Copie todo o conteúdo do arquivo [src/ui/api/app/.env.example](src/ui/api/app/.env.example) e cole em um novo arquivo chamado `.env` na mesma pasta ([src/ui/api/app/](src/ui/api/app/));
3. Execute o banco de dados e instância redis com o seguinte comando: `docker compose up database redis --build -d`

Pronto! O projeto está configurado. A partir de agora, toda vez que quiser iniciar o projeto basta executar o comando `go run main.go` dentro da pasta `src/ui/api/app`. Assim, o projeto estará disponível no endereço `http://localhost:8000`.

### **Para Testes de Qualidade (QA)**

Se você é um membro do time de qualidade (QA), siga os passos abaixo para executar as configurações apropriadas:

1. Copie todo o conteúdo do arquivo [src/ui/api/app/.env.example](src/ui/api/app/.env.example) e cole em um novo arquivo chamado `env-file` na mesma pasta ([src/ui/api/app/](src/ui/api/app/));

Pronto! O projeto está configurado. A partir de agora, toda vez que quiser iniciar o projeto basta executar o comando `docker compose up`. Assim, o projeto estará disponível no endereço `http://localhost:8000`.

Observações:
1. caso houve uma atualização no projeto e você não consegue ver essa atualização mesmo após executar `git pull`, execute o comando `docker compose up --build`.
2. sempre que for realizar os testes de qualidade para testar a aplicação, por favor, tenha certeza que você está no ramo `stage` do repositório.
