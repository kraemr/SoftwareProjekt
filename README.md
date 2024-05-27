# Xplocity lol

Eine React / Angular / Electron / hauptsache kein nodejs Web Applikation um Ausflugsziele in deiner Stadt zu finden

# Attractions Api
## delete attraction

send request with DELETE to
path: /api/attractions?id=10000
with the given attraction id you want deleted

## update attraction

send entire attraction json object in JSON PUT TO
path: /api/attractions

## add attraction ?

send entire attraction json object in JSON POST TO
path: /api/attractions

## get attractions filtered by city ?

GET
path: /api/attractions?city=Oppenheim
returns array of attraction Json Objects

## get attractions filtered by title ?

GET
path: /api/attractions?title=Schwimmbad
returns array of attraction Json Objects

## get attractions filtered by category ("type" in db)

GET
path: /api/attractions?category=Landmark
returns array of attraction Json Objects

## get single attraction by id

GET
path: /api/attractions?id=5
returns single attraction Json Object



# Notifications
Notification have info field and date field info contains html that should directly be added to a divs innerHTML 
```js
webSocket = new WebSocket("wss://"+document.location.host+"/notifications");
webSocket.onmessage = (event) => {
  console.log(event.data);
};
...
```


# Docker Compose

```bash
docker-compose up --build
```

# GO

In /app/api um zu Compilen

```bash
cd app/api/src
go build -o ../build/api main.go
```

# COMPOSE LOG/PRINTING

In Docker-Compose Output werden Print statements ohne newline nicht angezeigt!!!!!

# Testing Websocket Locally

In firefox you need to add an exception in certificates for localhost
