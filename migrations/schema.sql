CREATE DATABASE if not exists pie_test;
USE pie_test;

CREATE TABLE if not exists drivers (
       id VARCHAR(36) PRIMARY KEY,
       available_from DATETIME NOT NULL,
       available_till DATETIME NOT NULL
);

CREATE TABLE if not exists journeys (
       id VARCHAR(36) PRIMARY KEY,
       departure_time DATETIME NOT NULL,
       arrival_time DATETIME NOT NULL,
       departure_location VARCHAR(128) NOT NULL,
       arrival_location VARCHAR(128) NOT NULL
);

SET @LOC_1 = 'London';
SET @LOC_2 = 'Brighton';
SET @LOC_3 = 'Manchester';
SET @LOC_4 = 'Birmingham';

INSERT INTO drivers (id, available_from, available_till) VALUES
       ('DRIVER_1_ID', '2016-03-11 9:00:00.00', '2016-03-11 23:00:00.00'),
       ('DRIVER_2_ID', '2016-03-11 12:00:00.00', '2016-03-11 18:00:00.00'),
       ('DRIVER_3_ID', '2016-03-11 9:00:00.00', '2016-03-11 17:00:00.00');

INSERT INTO journeys (id, departure_location, arrival_location, departure_time, arrival_time) VALUES
       ('ROUTE_1_ID', @LOC_1, @LOC_2, '2016-03-11 15:00:00.00', '2016-03-11 17:00:00.00'),
       ('ROUTE_2_ID', @LOC_2, @LOC_1, '2016-03-11 18:00:00.00', '2016-03-11 20:00:00.00'),
       ('ROUTE_3_ID', @LOC_1, @LOC_3, '2016-03-11 09:00:00.00', '2016-03-11 13:00:00.00'),
       ('ROUTE_4_ID', @LOC_3, @LOC_1, '2016-03-11 14:00:00.00', '2016-03-11 18:00:00.00'),
       ('ROUTE_5_ID', @LOC_1, @LOC_4, '2016-03-11 12:00:00.00', '2016-03-11 15:00:00.00'),
       ('ROUTE_6_ID', @LOC_4, @LOC_1, '2016-03-11 16:00:00.00', '2016-03-11 19:00:00.00'),
       ('ROUTE_7_ID', @LOC_1, @LOC_2, '2016-03-11 06:00:00.00', '2016-03-11 08:00:00.00'),
       ('ROUTE_8_ID', @LOC_2, @LOC_1, '2016-03-11 09:00:00.00', '2016-03-11 11:00:00.00'),
       ('ROUTE_9_ID', @LOC_1, @LOC_3, '2016-03-11 22:00:00.00', '2016-03-12 02:00:00.00'),
       ('ROUTE_10_ID', @LOC_3, @LOC_1, '2016-03-12 03:00:00.00', '2016-03-12 07:00:00.00'),
       ('ROUTE_11_ID', @LOC_1, @LOC_4, '2016-03-11 20:00:00.00', '2016-03-11 23:00:00.00'),
       ('ROUTE_12_ID', @LOC_4, @LOC_1, '2016-03-11 23:45:00.00', '2016-03-12 02:45:00.00');
