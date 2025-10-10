# 🌤️ Weather Cloud Run

[![Go](https://img.shields.io/badge/Go-1.24-blue)](https://golang.org/)  
[![Cloud Run](https://img.shields.io/badge/Google%20Cloud-Cloud%20Run-orange)](https://cloud.google.com/run)  
[![Docker](https://img.shields.io/badge/Docker-Enabled-blue)](https://www.docker.com/)

Aplicação desenvolvida em **Go** para consulta de clima a partir de um **CEP válido**, retornando a temperatura atual em **Celsius, Fahrenheit e Kelvin**.  
O sistema está disponível no **Google Cloud Run**:  

👉 [Acessar aplicação online](https://weather-cloud-run-609455530745.southamerica-east1.run.app)

## 📌 Objetivo do projeto

Desenvolver um sistema em **Go** que receba um CEP, identifique a cidade correspondente e retorne o clima atual.  
O projeto contempla:

- Receber CEP válido de **8 dígitos**
- Buscar a localização via **ViaCEP**
- Consultar temperatura atual via **WeatherAPI**
- Converter a temperatura para **Celsius, Fahrenheit e Kelvin**
- Publicar no **Google Cloud Run**

## Requisitos do Projeto

* O sistema deve receber um CEP válido de 8 digitos
* O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
* O sistema deve responder adequadamente nos seguintes cenários:
  * Em caso de sucesso:
    * Código HTTP: **200**
    * Response Body: **{ "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }**
  * Em caso de falha, caso o CEP não seja válido (com formato correto):
    * Código HTTP: **422**
    * Mensagem: **invalid zipcode**
  * ​​​Em caso de falha, caso o CEP não seja encontrado:
    * Código HTTP: **404**
    * Mensagem: **can not find zipcode**
* Deverá ser realizado o deploy no Google Cloud Run.

### Dicas:

* Utilize a API **viaCEP** (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
* Utilize a API **WeatherAPI** (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
* Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: **F = C * 1,8 + 32**
* Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: **K = C + 273**
  * Sendo F = Fahrenheit
  * Sendo C = Celsius
  * Sendo K = Kelvin

### Entrega:

* O código-fonte completo da implementação.
* Testes automatizados demonstrando o funcionamento.
* Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
* Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.

## 📋 Requisitos do sistema

O sistema deve responder adequadamente em diferentes cenários:

### ✅ Sucesso

- **HTTP Code**: 200  
- **Response Body**:

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

### ❌ CEP inválido

- **HTTP Code: 422**
- **Response Body**:

```json
{
  "status_code": 422,
  "Message": "invalid zipcode"
}
```

### ❌ CEP não encontrado

- **HTTP Code: 404**
- **Response Body**:

```json
{
  "status_code": 404,
  "message": "can not find zipcode"
}
```

## 🚀 Como rodar o projeto localmente

### 1. Clonar repositório

```bash
git clone https://github.com/Berchon/weather-cloud-run.git
cd weather-cloud-run
```

### 2. Configurar variáveis de ambiente

Renomeie `.env.example` para `.env` e adicione sua chave da WeatherAPI:

```bash
cp .env.example .env
```

Exemplo `.env`:

```bash
WEATHER_API_KEY=sua_api_key_aqui
```

> ⚠️ A mesma chave deve ser configurada também no api/services.http para testes locais.

### 3. Rodar com Docker

#### Build da imagem Docker:

```bash
make build
```

#### Subir container:

```bash
make run
```

### 4. Rodar com Docker Compose

```bash
make up
```

### Para derrubar os containers:

```bash
make down
```

> ℹ️ make build cria a imagem Docker.

> ℹ️ make up usa docker-compose para subir a aplicação, tornando o passo mais simples para desenvolvimento local.

## 🧪 Testes de unidade

Para rodar os testes de unidade:

```bash
make test
```

O projeto inclui **tests unitários** e **mocking** para simular chamadas externas.

## ▶️ Executar  a aplicação

### Executar localmente

```bash
make run

echo -n "Retorna http status code 200: "; curl -s "http://localhost:8080/temperature/90040-000"

echo -n "Retorna http status code 404: "; curl -s "http://localhost:8080/temperature/90040999"

echo -n "Retorna http status code 422: "; curl -s "http://localhost:8080/temperature/1234567"
```

### Executar no google cloud run

```bash
echo -n "Retorna http status code 200: "; curl -s "https://weather-cloud-run-609455530745.southamerica-east1.run.app/temperature/90040-000"

echo -n "Retorna http status code 404: "; curl -s "https://weather-cloud-run-609455530745.southamerica-east1.run.app/temperature/90040999"

echo -n "Retorna http status code 422: "; curl -s "https://weather-cloud-run-609455530745.southamerica-east1.run.app/temperature/1234567"
```

#### Nota

> Também é possível testar endpoints manualmente via `api/api.http` e `api/services.http` usando REST Client no VSCode.

## 🔑 Tecnologias e APIs utilizadas

* Golang 1.24
* Docker / Docker Compose
* Google Cloud Run (Free Tier)
* ViaCEP API (https://viacep.com.br/)
* WeatherAPI (https://www.weatherapi.com/)

## ☁️ Deploy no Google Cloud Run

Aplicação publicada no **Google Cloud Run** (Free Tier):

👉 https://weather-cloud-run-609455530745.southamerica-east1.run.app

## 📂 Estrutura do projeto

```bash
weather-cloud-run/
├── .env.example            # Exemplo de variáveis de ambiente
├── .env                    # Variáveis de ambiente (não versionar)
├── Dockerfile              # Configuração do container
├── docker-compose.yaml     # Subir aplicação via Docker Compose
├── Makefile                # Comandos build/test/up/down
├── go.mod
├── go.sum
├── api/
│   ├── api.http            # Testes de endpoints
│   └── services.http       # Testes de serviços
├── cmd/
│   └── webserver/
│       └── main.go         # Entry point
├── internal/
│   ├── business/
│   │   ├── gateway/        # Protocolos (ViaCEP, Weather)
│   │   ├── model/          # Models e erros
│   │   └── usecase/        # Casos de uso (ex: get_temperature_by_zip_code)
│   └── infrastructure/
│       ├── configs/        # Configuração do ambiente
│       ├── dependencies/   # Injeção de dependências
│       ├── service/        # Serviços externos
│       └── webapp/         # HTTP handlers, request e routes

```

## 👨‍💻 Autor

Projeto desenvolvido por [Berchon](https://github.com/Berchon).