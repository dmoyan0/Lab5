Maquinas virtuales:
dist029-hUUWZwLHf2sn se debe ejecutar a:
Broker.go

dist099-6RJED2mGeXnR debe ejecutar a:
Comandante.go
Fulcrum1.go

dist031-jmsD482FZa5e debe ejecutar a:
Jeth.go
Fulcrum2.go

dist032-cG2d4v2aUEf5 debe ejecutar a:
Malkor.go
Fulcrum3.go

#COMENTARIO: En las MV vienen construidos los contenedores, por lo que solo se deberia escribir el comando de ejecución, si se quiere ejecutar desde 0 se deben borrar las imagenes creadas y seguir los siguientes pasos:
#PARA ELIMINAR LAS IMAGENES:
sudo docker stop {nombre-imagen} 
sudo docker rm {nombre-imagen}
ejemplo: 
sudo docker stop comandante-server
sudo docker rm comandante-server
Comandos Docker para la ejecución:
1.Broker
Entrar a la ruta LAB5/Lab5
Construcción -> sudo docker build --build-arg SERVER_TYPE=Broker -t broker-server .
Ejecución -> sudo docker run -d -p 50051:50051 --name broker-server -e SERVER_TYPE=Broker broker-server
sudo docker logs -f broker-server (para ver logs del broker)

2.Fulcrum1
Entrar a la ruta Lab5
Construcción -> sudo docker build --build-arg SERVER_TYPE=Fulcrum1 -t fulcrum1-server .
Ejecución -> sudo docker run -d -p 60051:60051 --name fulcrum1-server -e SERVER_TYPE=Fulcrum1 fulcrum1-server
sudo docker logs -f fulcrum1-server (para ver logs del fulcrum)

3.Fulcrum2
Entrar a la ruta LAB5/Lab5
Construcción -> sudo docker build --build-arg SERVER_TYPE=Fulcrum2 -t fulcrum2-server .
Ejecución -> sudo docker run -d -p 60052:60052 --name fulcrum2-server -e SERVER_TYPE=Fulcrum2 fulcrum2-server
sudo docker logs -f fulcrum2-server (para ver logs del fulcrum)

4.Fulcrum3
Entrar a la ruta LAB5/Lab5
Construcción -> sudo docker build --build-arg SERVER_TYPE=Fulcrum3 -t fulcrum3-server .
Ejecución -> sudo docker run -d -p 60053:60053 --name fulcrum3-server -e SERVER_TYPE=Fulcrum3 fulcrum3-server
sudo docker logs -f fulcrum3-server (para ver logs del fulcrum)

5.Jeth 
Entrar a la ruta LAB5/Lab5
Construcción->sudo docker build --build-arg SERVER_TYPE=Jeth -t jeth-server .
Ejecución->sudo docker run -d -p 60055:60055 --name jeth-server -e SERVER_TYPE=Jeth jeth-server
sudo docker exec -it jeth-server /bin/bash
entrar a la ruta MV3
escribir: make run2

6. Malkor
Entrar a la ruta LAB5/Lab5
Construcción->sudo docker build --build-arg SERVER_TYPE=Malkor -t malkor-server .
Ejecución->sudo docker run -d -p 60054:60054 --name malkor-server -e SERVER_TYPE=Malkor malkor-server
sudo docker exec -it malkor-server /bin/bash
entrar a la ruta MV4
escribir: make run2

7. Comandante
Entrar a la ruta Lab5
Construcción->sudo docker build --build-arg SERVER_TYPE=Comandante -t comandante-server .
Ejecución->sudo docker run -d -p 50052:50052 --name comandante-server -e SERVER_TYPE=Comandante comandante-server
sudo docker exec -it comandante-server /bin/bash
entrar a la ruta MV2
escribir: make run2

