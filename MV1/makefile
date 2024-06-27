# Define los nombres de los binarios
BIN1 := Broker

# Define los archivos fuente
SRC1 := Broker.go

# Target para compilar el primer archivo
$(BIN1): $(SRC1)
	go build -o $(BIN1) $(SRC1)

# Target para ejecutar el primer archivo
run: $(BIN1)
	./$(BIN1)

# Limpieza de binarios
clean:
	rm -f $(BIN1) $(BIN2) *.txt
