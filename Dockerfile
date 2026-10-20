# Imagem da API. Quem usa é o docker-compose.yml: "docker compose up -d --build"
# sobe o banco e a API sem precisar de Go instalado na máquina.

# Etapa 1: compila a API com o Go oficial. Só esta etapa tem o Go.
FROM golang:1.27-alpine AS build
WORKDIR /src

# Baixa as bibliotecas antes de copiar o código: se só o código mudar, o Docker
# reaproveita esta etapa e o build fica bem mais rápido.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# CGO_ENABLED=0 gera um executável que não depende de nada do sistema.
# A tag timetzdata embute os fusos horários no executável, para o vendidaEm
# sair no horário de Brasília (-03:00) mesmo numa imagem mínima.
RUN CGO_ENABLED=0 go build -tags timetzdata -o /syncvault .

# Etapa 2: a imagem final tem só o executável e as migrations, que a API aplica
# quando sobe. Fica com poucos MB, sem o Go nem o código-fonte.
FROM alpine:3.22
WORKDIR /app
COPY --from=build /syncvault ./syncvault
COPY migrations ./migrations

# Horário de Brasília, e o Gin sem as mensagens de depuração.
ENV TZ=America/Sao_Paulo
ENV GIN_MODE=release

EXPOSE 8080
CMD ["./syncvault"]
