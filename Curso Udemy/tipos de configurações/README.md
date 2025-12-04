## Tipos de configurações em projetos 

### 1- .env -> para váriaveis de ambiente
para ativar as variaveis de ambiente no linux use:
`source .env`
é necessario para que a aplicação encontre as váriaveis
caso não faça isso pode ocorrer o erro de "varible not found" ou algo assim
e no .env precisa dar `export` na frente da variável pra funcionar

Para o windows da pra configurar pelo powershell
$env:NOME_VARIAVEL=8000
porém é de modo temporário igual o `source .env` que usamos no linux

Podemos usar também o pacode [godotenv](https://github.com/joho/godotenv) conforme mostrado em [aqui](/env/main.go)
que elimina as diferenças de SO

### 2- por flags
exemplo:
`go run main.go -porta 5000`
para configurar a porta

### Usando Json