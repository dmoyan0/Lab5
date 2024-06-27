Comandos Docker

1.Broker
Construcción -> docker build --build-arg SERVER_TYPE=Broker -t broker-server .
Ejecución -> docker run -d -p 50051:50051 --name broker-server -e SERVER_TYPE=Broker broker-server

2.Fulcrum1
Construcción -> docker build --build-arg SERVER_TYPE=Fulcrum1 -t fulcrum1-server .
Ejecución -> docker run -d -p 60051:60051 --name fulcrum1-server -e SERVER_TYPE=Fulcrum1 fulcrum1-server

3.Fulcrum2
Construcción -> docker build --build-arg SERVER_TYPE=Fulcrum2 -t fulcrum2-server .
Ejecución -> docker run -d -p 60052:60052 --name fulcrum2-server -e SERVER_TYPE=Fulcrum2 fulcrum2-server

4.Fulcrum3
Construcción -> docker build --build-arg SERVER_TYPE=Fulcrum3 -t fulcrum3-server .
Ejecución -> docker run -d -p 60053:60053 --name fulcrum3-server -e SERVER_TYPE=Fulcrum3 fulcrum3-server

5.Jeth 
Construcción->docker build --build-arg SERVER_TYPE=Jeth -t jeth-server .
Ejecución->docker run -d -p 60055:60055 --name jeth-server -e SERVER_TYPE=Jeth jeth-server

6. Malkor
Construcción->docker build --build-arg SERVER_TYPE=Malkor -t malkor-server .
Ejecución->docker run -d -p 60054:60054 --name malkor-server -e SERVER_TYPE=Malkor malkor-server

7. Comandante
Construcción->docker build --build-arg SERVER_TYPE=Comandante -t comandante-server .
Ejecución->docker run -d -p 50052:50052 --name comandante-server -e SERVER_TYPE=Comandante comandante-server


