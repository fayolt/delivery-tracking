# Delivery Tracking

### Database

* Postgres 9.6 

* Database structure available in `delivery-tracking.sql`

### Assumptions

* Job is added back to queue for later processing after failure
* Job processing failure related only to network/db unavailability/external issues
* Data doesn't violate any integrity constraints 