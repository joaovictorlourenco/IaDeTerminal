![2025-03-30-12-13-58](https://github.com/user-attachments/assets/ba8db41d-c06f-4429-8fe3-05273623a7e0)

# Groot Terminal

Groot Terminal é uma IA integrada ao terminal desenvolvida em Go. Este projeto oferece uma interface interativa que combina a eficiência do terminal com funcionalidades inteligentes, facilitando a interação do usuário com comandos e operações através de uma interface mais amigável e colorida.

## Características

- **Integração com IA:** Respostas inteligentes e interativas diretamente no terminal.
- **Interface Visual:** Utilização de spinners e textos coloridos para melhorar a experiência do usuário.
- **Gerenciamento de Configurações:** Suporte ao carregamento de variáveis de ambiente por meio de arquivos `.env`.
- **CLI Robusto:** Estrutura baseada em comandos com suporte completo para a criação de uma interface de linha de comando.

## Tecnologias e Bibliotecas Utilizadas

O projeto é escrito em Go e faz uso das seguintes bibliotecas:

- **[github.com/briandowns/spinner](https://github.com/briandowns/spinner) v1.23.2:** Exibe animações tipo spinner no terminal, aprimorando a experiência visual.
- **[github.com/fatih/color](https://github.com/fatih/color) v1.18.0:** Permite a utilização de cores no output do terminal para destacar informações.
- **[github.com/joho/godotenv](https://github.com/joho/godotenv) v1.5.1:** Carrega variáveis de ambiente a partir de um arquivo `.env`, facilitando a configuração do projeto.
- **[github.com/spf13/cobra](https://github.com/spf13/cobra) v1.9.1:** Framework para a criação de interfaces de linha de comando (CLI) robustas e intuitivas.

## Instalação

1. **Pré-requisitos:**  
   Certifique-se de ter o Go instalado (versão 1.23.0 ou superior).

2. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/groot-terminal.git
   ```
3. **Navegue até o diretório do projeto:**

   ```bash
   cd groot-terminal
   ```

4. **Instale as dependências:**

   ```bash
   go mod download
   ```

5. **Compile o projeto:**
   ```bash
   go build -o groot-terminal
   ```

## Uso

Após compilar, execute o terminal:

```bash
./groot-terminal
```

ou

```bash
go run main.go
```

O aplicativo iniciará e você poderá interagir com a IA integrada, explorando os diversos comandos e funcionalidades.

## Configuração

O projeto utiliza um arquivo `.env` para configurar variáveis de ambiente. Crie um arquivo `.env` na raiz do projeto com as configurações necessárias, por exemplo:

```dotenv
API_KEY=sua_chave_de_api
OUTRA_CONFIG=valor
```

## Contribuição

Contribuições são bem-vindas!  
Se você deseja ajudar a melhorar o Groot Terminal, sinta-se à vontade para abrir issues e enviar pull requests.
