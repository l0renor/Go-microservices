# Microservice Cinema
## Start
Einfach aus dem Projekt-Root `sudo docker-compose up --abort-on-container-exit --exit-code-from client` aufrufen, dann läuft der Client-Test komplett durch.
Falls es neue Commits gab, die Einfluss auf die Images haben, zuvor `./reload_images.sh` aufrufen, um die neuen Images vom Jenkins zu holen.