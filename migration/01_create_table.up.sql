CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  

SELECT uuid_generate_v4();  

CREATE TABLE users(
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    first_name VARCHAR,
    last_name VARCHAR,
    phone_number VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users(first_name,last_name,phone_number) VALUES('Akbar','Komilov','+998 11 775 55 45');
INSERT INTO users(first_name,last_name,phone_number) VALUES('John','Doe','+998 11 888 55 45');
INSERT INTO users(first_name,last_name,phone_number) VALUES('Ali','Usmonov','+998 11 555 55 45');
INSERT INTO users(first_name,last_name,phone_number) VALUES('Temur','Husanov','+998 11 889 55 55');
INSERT INTO users(first_name,last_name,phone_number) VALUES('Rustam','Olimov','+998 11 222 55 22');