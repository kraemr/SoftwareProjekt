CREATE DATABASE IF NOT EXISTS SITE_DB;
USE SITE_DB;
CREATE TABLE IF NOT EXISTS USER (id SERIAL PRIMARY KEY,email TEXT,password TEXT);
DROP TABLE ATTRACTION_ENTRY;
CREATE TABLE IF NOT EXISTS ATTRACTION_ENTRY(id SERIAL PRIMARY KEY,title varchar(64),
type varchar(32),recommended_count int,city Text,info Text,approved BOOLEAN,PosX Float,PosY Float,stars Float);


-- Insert sample data into the ATTRACTION_ENTRY table for places in Germany
INSERT INTO ATTRACTION_ENTRY (title, type, recommended_count, city, info, approved, PosX, PosY, stars)
VALUES
('Brandenburg Gate', 'Landmark', 12000, 'Berlin', 'An 18th-century neoclassical monument in Berlin, one of the best-known landmarks of Germany.', TRUE, 52.5163, 13.3777, 4.8),
('Neuschwanstein Castle', 'Castle', 15000, 'Schwangau', 'A 19th-century historicist palace on a rugged hill above the village of Hohenschwangau near Füssen in southwest Bavaria.', TRUE, 47.5576, 10.7498, 4.9),
('Cologne Cathedral', 'Cathedral', 18000, 'Cologne', 'A renowned monument of German Catholicism and Gothic architecture, also a World Heritage Site.', TRUE, 50.9413, 6.9583, 4.8),
('Heidelberg Castle', 'Castle', 8000, 'Heidelberg', 'A famous ruin in Germany and landmark of Heidelberg, built on a hill overlooking the city.', TRUE, 49.4106, 8.7158, 4.7),
('Miniatur Wunderland', 'Museum', 9000, 'Hamburg', 'The world\'s largest model railway exhibition, located in Hamburg.', TRUE, 53.5430, 9.9881, 4.9),
('Zugspitze', 'Mountain', 6000, 'Garmisch-Partenkirchen', 'The highest peak of the Wetterstein Mountains as well as the highest mountain in Germany.', TRUE, 47.4211, 10.9853, 4.8),
('Marienplatz', 'Square', 11000, 'Munich', 'A central square in the city center of Munich, featuring the New Town Hall and the Glockenspiel.', TRUE, 48.1374, 11.5755, 4.7),
('Sanssouci Palace', 'Palace', 7000, 'Potsdam', 'The former summer palace of Frederick the Great, King of Prussia, in Potsdam, near Berlin.', TRUE, 52.4011, 13.0416, 4.7),
('AMONGUS Palace', 'Palace', 7000, 'Imposter', 'The former summer palace of Frederick the SUS, King of SUSSEX, in SUSSEX.', FALSE, 12.4011, 13.0416, 0);