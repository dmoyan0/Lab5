# Define los nombres de los binarios
BIN1 := Fulcrum2
BIN2 := Jeth

# Define los archivos fuente
SRC1 := Fulcrum2.go
SRC2 := Jeth.go

# Target para compilar el primer archivo
$(BIN1): $(SRC1)
	go build -o $(BIN1) $(SRC1)

# Target para compilar el segundo archivo
$(BIN2): $(SRC2)
	go build -o $(BIN2) $(SRC2)

# Target para ejecutar el primer archivo
run1: $(BIN1)
	./$(BIN1)

# Target para ejecutar el segundo archivo
run2: $(BIN2)
	./$(BIN2)

# Target para ejecutar ambos archivos en orden
run_all: run1 run2

# Target por defecto para compilar y ejecutar ambos archivos
all: $(BIN1) $(BIN2) run_all

# Limpieza de binarios
clean:
	rm -f $(BIN1) $(BIN2) *.txt
