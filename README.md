# **Avenue - Upload File**
> A aplicação realiza o upload e donwload de arquivos 
## Instalação e execução da aplicação
-----
### Testes

Para executar a stack de testes basta executar o seguinte comando:

```sh
go test ./...
```

### Executando local

Para executar o projeto localmente basta executar o seguinte comando:

```sh
go run src/main.go
```

---
### Running the Project

Collection no Postman podem ser encontradas aqui:
<br>

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/0ea80040988a7f050a9a?action=collection%2Fimport)


---
## Configurações

> Lista com todas as configurações possíveis e seus defaults caso existam

- **PORT**: Porta na qual o servidor web ficará disponível (valor default: 4545);
- **UPLOAD**: Realiza o upload do arquivo na memoria ou no file system (valor: "fs" ou "memory");
