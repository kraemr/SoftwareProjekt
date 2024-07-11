# Xplocity lol

## Added Features since 04.07.2024
+better and more detailed route planer<br>
+added a review tab<br>
+added favorite count<br>
+added favourite / unfavourite button<br>
+better favorite management<br>
+added reviews<br>
+make reviews owner-editable<br>
+added star rating based on average of all reviews<br>
+added user moderation panel<br>
+added seeing reviews for attractions for moderating<br>
+added seeing reviews from users for moderating<br>
+a bunch of bug fixes<br>

A custom deployable package using docker for a go backend and html frontend webapp to explore, add and review special places in your city!
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

# Testing:

**Important** When changing the test files, the Docker must always be re-composed!

```bash
docker exec -it softwareprojekt-webserver-1 bash

cd tests/

go test -v 
```

alternative: just use this command when composing (Sadly, only found out about it at the end...)

```bash
docker-compose up --build --watch
```
