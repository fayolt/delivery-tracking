CREATE TABLE drivers (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);

INSERT INTO drivers (name) VALUES ('Phil');
INSERT INTO drivers (name) VALUES ('Nicholas');
INSERT INTO drivers (name) VALUES ('William');

CREATE TABLE locations (
  id SERIAL PRIMARY KEY,
  longitude INT NOT NULL,
  latitude INT NOT NULL,
  driver_id INT REFERENCES drivers(id)
);
