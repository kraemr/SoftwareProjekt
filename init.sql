DROP DATABASE SITE_DB;
CREATE DATABASE IF NOT EXISTS SITE_DB;
USE SITE_DB;

CREATE TABLE CITY_MODERATOR(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email TEXT,
    password TEXT,
    city TEXT,
    username TEXT
);



CREATE TABLE IF NOT EXISTS USER (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email TEXT,
    password TEXT,
    city TEXT,
    username TEXT,
    active BOOLEAN
);
-- create index
CREATE INDEX _USER_INDEX ON USER(id);


CREATE TABLE IF NOT EXISTS USER_PREFERENCES (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    periodic_recommendations BOOLEAN
);

-- create index
CREATE INDEX _USER_PREFS_INDEX ON USER_PREFERENCES(user_id);

CREATE TABLE USER_FAVORITE(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    attraction_id INT NOT NULL UNIQUE
);
CREATE INDEX _USER_FAVORITE_INDEX ON USER_FAVORITE(user_id);

CREATE TABLE IF NOT EXISTS ATTRACTION_ENTRY(
    id INT AUTO_INCREMENT PRIMARY KEY,
    title varchar(64),
    type varchar(32),
    recommended_count int,
    city Text,
    street Text,
    housenumber Text,
    info Text,
    approved BOOLEAN,
    PosX Float,
    PosY Float,
    stars Float,
    img_url TEXT,
    added_by INT
);

CREATE INDEX _ATTRACTION_ENTRY_INDEX ON ATTRACTION_ENTRY(city);


-- To get Username JOIN
-- TODO Add DATE
CREATE TABLE IF NOT EXISTS ATTRACTION_REVIEW (
    id BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,   
    user_id INT,
    attraction_id INT,
    text TEXT,
    stars FLOAT,
    date DATE
);
CREATE INDEX _ATTRACTION_REVIEW_INDEX ON ATTRACTION_REVIEW(attraction_id);




CREATE TABLE USER_NOTIFICATIONS(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    info TEXT NOT NULL,
    date DATE
);
CREATE INDEX _USER_NOTIFICATIONS_INDEX ON USER_NOTIFICATIONS(user_id);



CREATE TABLE CITY_NOTIFICATIONS(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    info TEXT NOT NULL,
    date DATE,
    city TEXT
);
CREATE INDEX _CITY_NOTIFICATIONS_INDEX ON USER_NOTIFICATIONS(user_id);


-- TESTDATA
-- passwort1234
INSERT INTO USER(id,email,password,city,username,active) VALUES(911111,"admin@testemail.com","$argon2id$v=19$m=2048,t=1,p=2$m0Ro6ArcaMfanzBFGVmQCw$vmDrLnu2CfevEJwJh/KeVu53cScOfjYzF57jNIFPJ4Q","Oppenheim","adminman",TRUE);
INSERT INTO USER(id,email,password,city,username,active) VALUES(911112,"test@testemail.com","$argon2id$v=19$m=2048,t=1,p=2$m0Ro6ArcaMfanzBFGVmQCw$vmDrLnu2CfevEJwJh/KeVu53cScOfjYzF57jNIFPJ4Q","Müllhausen","testman",TRUE);
INSERT INTO USER(id,email,password,city,username,active) VALUES(911113,"meeenz@meeenz.com","$argon2id$v=19$m=2048,t=1,p=2$m0Ro6ArcaMfanzBFGVmQCw$vmDrLnu2CfevEJwJh/KeVu53cScOfjYzF57jNIFPJ4Q","Mainz","meeenzman",TRUE);
-- passwort1234
INSERT INTO CITY_MODERATOR(id,email,password,city,username) VALUES(911113,"meeenz@meeenz.com","$argon2id$v=19$m=2048,t=1,p=2$m0Ro6ArcaMfanzBFGVmQCw$vmDrLnu2CfevEJwJh/KeVu53cScOfjYzF57jNIFPJ4Q","Mainz","meeenzman");


-- Review
INSERT INTO ATTRACTION_ENTRY (id,title, type, recommended_count, city, street, housenumber, info, approved, PosX, PosY, stars, img_url,added_by) VALUES
(1,'Brandenburg Gate', 'Monument', 12000, 'Berlin', 'Pariser Platz', '1', 'An 18th-century neoclassical monument in Berlin, one of the most well-known landmarks of Germany.', TRUE, 52.5163, 13.3777, 4.7, 'https://example.com/brandenburg_gate.jpg',911111),
(2,'Neuschwanstein Castle', 'Castle', 11000, 'Schwangau', 'Neuschwansteinstraße', '20', 'A 19th-century Romanesque Revival palace on a rugged hill above the village of Hohenschwangau near Füssen in southwest Bavaria.', TRUE, 47.5576, 10.7498, 4.8, 'https://example.com/neuschwanstein_castle.jpg',911111),
(3,'Cologne Cathedral', 'Cathedral', 13000, 'Cologne', 'Domkloster', '4', 'A renowned monument of German Catholicism and Gothic architecture and is a World Heritage Site.', TRUE, 50.9413, 6.9583, 4.9, 'https://example.com/cologne_cathedral.jpg',911111),
(4,'Heidelberg Castle', 'Castle', 9500, 'Heidelberg', 'Schlosshof', '1', 'A famous ruin in Germany and landmark of Heidelberg.', TRUE, 49.4106, 8.7153, 4.6, 'https://example.com/heidelberg_castle.jpg',911111),
(5,'Sanssouci Palace', 'Palace', 8000, 'Potsdam', 'Maulbeerallee', '3', 'The former summer palace of Frederick the Great, King of Prussia, in Potsdam, near Berlin.', TRUE, 52.4044, 13.0388, 4.7, 'https://example.com/sanssouci_palace.jpg',911111),
(6,'Miniatur Wunderland', 'Museum', 15000, 'Hamburg', 'Kehrwieder', '2', 'A model railway attraction in Hamburg, and the largest of its kind in the world.', TRUE, 53.5436, 9.9886, 4.8, 'https://example.com/miniatur_wunderland.jpg',911111),
(9000,'Zugspitze', 'Mountain', 7000, 'Garmisch-Partenkirchen', '', '', 'The highest peak of the Wetterstein Mountains as well as the highest mountain in Germany.', TRUE, 47.4210, 10.9850, 4.9, 'https://example.com/zugspitze.jpg',911111),
(911113,'English Garden', 'Park', 6000, 'Munich', 'Englischer Garten', '', 'A large public park in the centre of Munich, stretching from the city center to the northeastern city limits.', FALSE, 48.1584, 11.5944, 4.7, 'https://example.com/english_garden.jpg',911111);


INSERT INTO ATTRACTION_REVIEW(user_id,attraction_id,text,stars) VALUES(911113,9000,"Great place for my trad wife and 50 kids would go again, was able to drink my beer in peace without my wife nagging",5);

INSERT INTO USER_NOTIFICATIONS(user_id,info,date) VALUES(911111,"<p> GOODBYE WORLD </p>","2000-01-01");
INSERT INTO CITY_NOTIFICATIONS(info,date,city)  VALUES("Kostenlose Döner","2000-01-01","Oppenheim");
-- Attraction

INSERT INTO USER_FAVORITE(user_id,attraction_id) VALUES(911111,9000); 

-- add foreign key constraints
ALTER TABLE ATTRACTION_REVIEW ADD CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);
ALTER TABLE ATTRACTION_REVIEW ADD CONSTRAINT attraction_id_fk FOREIGN KEY (attraction_id) REFERENCES ATTRACTION_ENTRY(id);
ALTER TABLE USER_NOTIFICATIONS ADD CONSTRAINT notif_user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);
ALTER TABLE USER_PREFERENCES ADD CONSTRAINT pref_user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);
ALTER TABLE USER_FAVORITE ADD CONSTRAINT fav_user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);
ALTER TABLE USER_FAVORITE ADD CONSTRAINT fav_attraction_id_fk FOREIGN KEY (attraction_id) REFERENCES ATTRACTION_ENTRY(id);




-- Review
-- TESTDATA
