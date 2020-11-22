# Delivery Tracking

### Database

* Postgres 9.6 

* Database structure available in `database/fixture.sql`

### Starting the backend app

* Prerequisites

    `docker` & `docker compose`

* Provide needed environment variables

    Add a `.env` file in the root directory of the project containing the environment variables

* Default `.env` file content

```
PORT=5000
MAX_PROCESSORS=5
POSTGRES_USER=postgres
POSTGRES_DB=delivery-tracking
POSTGRES_HOST_AUTH_METHOD=trust
```

* Build the docker images

```
docker-compose build
```
* Start the docker containers 

```
docker-compose up
```

* Stop the docker containers 

```sh
docker-compose down
```

Now the backend app should be available at [`localhost:5000`](http://localhost:5000)

### Available endpoints

* Send a driver's location `POST` [`/api/v1/locations`](http://localhost:5000/api/v1/locations)
* Retrieve all locations `GET` [`/api/v1/locations`](http://localhost:5000/api/v1/locations)

### Assumptions

* Job is added back to queue for later processing after failure
* Job processing failure related only to network/db unavailability/external issues
* Data doesn't violate any integrity constraints

### Possible Improvments

* Drain job queue and save all available locations before shutting down the server
* 