#!/bin/bash

# Configurações do banco de dados
CONTAINER_NAME="hexagonal_postgres"
POSTGRES_USER="hexagonal_user"
POSTGRES_PASSWORD="hexagonal_password"
POSTGRES_DB="hexagonal_db"
POSTGRES_PORT="5432"

# Verifica se o Docker está instalado
if ! command -v docker &> /dev/null
then
    echo "Docker não está instalado. Instale o Docker e tente novamente."
    exit 1
fi

echo "Iniciando o container Docker com PostgreSQL..."

# Remove qualquer container com o mesmo nome
if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
    echo "Removendo container antigo com o nome $CONTAINER_NAME..."
    docker rm -f $CONTAINER_NAME
fi

# Inicia o container PostgreSQL
docker run --name $CONTAINER_NAME \
    -e POSTGRES_USER=$POSTGRES_USER \
    -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
    -e POSTGRES_DB=$POSTGRES_DB \
    -p $POSTGRES_PORT:5432 \
    -d postgres:latest

if [ $? -ne 0 ]; then
    echo "Falha ao iniciar o container Docker."
    exit 1
fi

echo "Container $CONTAINER_NAME iniciado com sucesso!"

# Espera o banco de dados estar pronto
echo "Aguardando inicialização do PostgreSQL..."
sleep 5

# Cria conexão no banco e testa a conexão
echo "Testando a conexão com o banco de dados..."
docker exec -i $CONTAINER_NAME psql -U $POSTGRES_USER -d $POSTGRES_DB -c "\l"

if [ $? -eq 0 ]; then
    echo "Banco de dados configurado e conectado com sucesso!"
    echo "Host: localhost"
    echo "Porta: $POSTGRES_PORT"
    echo "Usuário: $POSTGRES_USER"
    echo "Senha: $POSTGRES_PASSWORD"
    echo "Banco: $POSTGRES_DB"
else
    echo "Falha ao conectar no banco de dados."
fi
