# Usa la imagen base especificada
ARG BASE_IMAGE=golang:1.22.1
FROM ${BASE_IMAGE} AS builder

# Define los puertos de los servidores
ARG BROKER_PORT=50051
ARG COMANDANTE_PORT=50052
ARG FULCRUM1_PORT=60051
ARG FULCRUM2_PORT=60052
ARG FULCRUM3_PORT=60053
ARG MALKOR_PORT=60054
ARG JETH_PORT=60055

ARG SERVER_TYPE

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum del directorio raíz al contenedor
COPY go.mod .
COPY go.sum .

# Descarga e instala las dependencias de Go
RUN go mod download

# Copia el resto del código de la aplicación al contenedor
COPY . .

# Establece la variable de entorno SERVER_TYPE
ENV SERVER_TYPE=${SERVER_TYPE}

# Comando para compilar y ejecutar el servidor basado en SERVER_TYPE
CMD if [ "$SERVER_TYPE" = "Broker" ]; then \
        PORT=$BROKER_PORT; \
        cd /app/MV1; \
        go build -o broker-server Broker.go; \
        ./broker-server; \
    elif [ "$SERVER_TYPE" = "Comandante" ]; then \
        PORT=$COMANDANTE_PORT; \
        cd /app/MV2; \
        go build -o comandante-server Comandante.go; \
        ./comandante-server; \
    elif [ "$SERVER_TYPE" = "Fulcrum1" ]; then \
        PORT=$FULCRUM1_PORT; \
        cd /app/MV2; \
        go build -o fulcrum1-server Fulcrum1.go; \
        ./fulcrum1-server; \
    elif [ "$SERVER_TYPE" = "Fulcrum2" ]; then \
        PORT=$FULCRUM2_PORT; \
        cd /app/MV3; \
        go build -o fulcrum2-server Fulcrum2.go; \
        ./fulcrum2-server; \
    elif [ "$SERVER_TYPE" = "Fulcrum3" ]; then \
        PORT=$FULCRUM3_PORT; \
        cd /app/MV4; \
        go build -o fulcrum3-server Fulcrum3.go; \
        ./fulcrum3-server; \
    elif [ "$SERVER_TYPE" = "Malkor" ]; then \
        PORT=$MALKOR_PORT; \
        cd /app/MV4; \
        go build -o malkor-server Malkor.go; \
        ./malkor-server; \
    elif [ "$SERVER_TYPE" = "Jeth" ]; then \
        PORT=$JETH_PORT; \
        cd /app/MV3; \
        go build -o jeth-server Jeth.go; \
        ./jeth-server; \
    else \
        echo "Invalid SERVER_TYPE argument."; \
    fi
