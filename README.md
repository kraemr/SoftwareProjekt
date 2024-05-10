# Xplocity lol

Eine React / Angular / Electron / hauptsache kein nodejs Web Applikation um Ausflugsziele in deiner Stadt zu finden

# Docker Compose

´´´bash
docker-compose up --build
´´´

# GO

In /app/api um zu Compilen

´´´bash
cd app/api/src
go build -o ../build/api main.go
´´´

# Einträge
```json
Title: Europa Park
Type: Family Trip?
Recommended By: 100 // 100 Users like it there
City: Kaiserslautern
Info: string with info about the attraction
PosX: 49.0192190123°
PosY: 8.01293123°
```
Einträge können anhand der PosX und PosY geladen werden als Marker auf der Map.(Mit Leaflet?)

Info ist dann ein string mit Informationen zu der Attraktion

Recommended by zeigt wie viele Nutzer eine gute Erfahrung dort hatten

Type oder Genre oder wie auch immer wir es dann nennen ist später zum Filtern da.


## Einträge in DB
Einträge könnten wir theoretisch cachen, denn das problem ist das die recht groß sind in der DB.
Zusätzlich könnten wir simple Kompressionsalgos benutzen wie RLE, Huffman,Lz4 ... um die Größe zu reduzieren.

```sql
CREATE TABLE attraction_entry(id AUTO_INCREMENT INTEGER PRIMARY KEY,title varchar(64),type varchar(32),recommended_count int,city Text,info Text,PosX double,PosY double)
```

# COMPOSE LOG/PRINTING
In Docker-Compose Output werden Print statements ohne newline nicht angezeigt!!!!!
