# Multiplex-Kino Protocol
![Service Diagramm](doc/Communications.png)
## RPCs
### reservation_service
- CreateReservation (userID, screeningID, nrOfSteats) returns (reservationID) 
- ActivateReservation (reservationID) returns ()
- DeleteReservation (reservationID) returns ()
- DeleteReservationsWithScreening (screeningID) returns ()
- GetReservation (reservationID) returns (userID, screeningID, nrOfSeats, active)
- GetReservations () returns ([](userID, screeningID, nrOfSeats, active))

### user_service
- CreateUser (name) returns (userID)
- DeleteUser (userID) returns ()
- GetUser (userID) returns (name)
- GetUsers () returns ([](userID, name))
- AddUserReservation (userID, reservationID) returns ()
- DeleteUserReservation (userID, reservationID) returns ()

### screening_service
- CreateScreening (movieID,roomID) returns (screeningID)
- ChangeFreeSeats (screeningID,change) returns ()
- DeleteScreening (screeningID) returns () 
- DeleteScreeningsWithRoom (roomID) returns ()
- DeleteScreeningsWithMovie (movieID) returns ()
- GetScreening (screeningID) returns (movieID, roomID ,nrOfFreeSeats)
- GetScreenings () returns ([](movieID, roomID, nrOfFreeSeats))

### movie_service
- CreateMovie(name) returns (movieID)
- DeleteMovie(movieID) returns ()
- GetMovie(movieID) returns (title)
- GetMovies() returns ([](title, movieID))

### room_service
- CreateRoom(name, nrOfSeats) returns (roomID)
- DeleteRoom(roomID) returns ()
- GetRoom(roomID) returns (name, roomID, nrOfSeats)
- GetRooms() returns ([](name, roomID, nrOfSeats))