DROP DATABASE SITE_DB;
CREATE DATABASE IF NOT EXISTS SITE_DB;
USE SITE_DB;

CREATE TABLE IF NOT EXISTS USER (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email TEXT,
    password TEXT,
    city TEXT,
    username TEXT,
    admin BOOLEAN
);

CREATE TABLE IF NOT EXISTS USER_PREFERENCES (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    periodic_recommendations BOOLEAN
);

CREATE TABLE USER_FAVORITE(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    attraction_id INT NOT NULL,
    type varchar(32) NOT NULL,
    city varchar(32) NOT NULL
);

CREATE TABLE IF NOT EXISTS ATTRACTION_ENTRY(
    id INT AUTO_INCREMENT PRIMARY KEY,title varchar(64),
    type varchar(32),
    recommended_count int,
    city Text,
    info Text,
    approved BOOLEAN,
    PosX Float,
    PosY Float,
    stars Float
);

-- To get Username JOIN
-- TODO Add DATE
CREATE TABLE IF NOT EXISTS ATTRACTION_REVIEW (
    id SERIAL PRIMARY KEY,   
    user_id INT,
    attraction_id INT,
    text TEXT,
    stars FLOAT
);


CREATE TABLE USER_NOTIFICATIONS(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    info TEXT NOT NULL,
    date DATE
);

-- add foreign key constraints
ALTER TABLE ATTRACTION_REVIEW ADD CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);
ALTER TABLE ATTRACTION_REVIEW ADD CONSTRAINT attraction_id_fk FOREIGN KEY (attraction_id) REFERENCES ATTRACTION_ENTRY(id);
ALTER TABLE USER_NOTIFICATIONS ADD CONSTRAINT notif_user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);
ALTER TABLE USER_PREFERENCES ADD CONSTRAINT pref_user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);

ALTER TABLE USER_FAVORITE ADD CONSTRAINT fav_user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);
ALTER TABLE USER_FAVORITE ADD CONSTRAINT fav_attraction_id_fk FOREIGN KEY (attraction_id) REFERENCES ATTRACTION_ENTRY(id);



-- TESTDATA
--- Users
-- ALL users have the password: passwort1234
-- Obviously delete em later ...
INSERT INTO USER(id,email,password,city,username,admin) VALUES(911111,"admin@testemail.com","$argon2id$v=19$m=2048,t=1,p=2$m0Ro6ArcaMfanzBFGVmQCw$vmDrLnu2CfevEJwJh/KeVu53cScOfjYzF57jNIFPJ4Q","Oppenheim","adminman",TRUE);
INSERT INTO USER(id,email,password,city,username,admin) VALUES(911112,"test@testemail.com","$argon2id$v=19$m=2048,t=1,p=2$m0Ro6ArcaMfanzBFGVmQCw$vmDrLnu2CfevEJwJh/KeVu53cScOfjYzF57jNIFPJ4Q","Müllhausen","testman",FALSE);
INSERT INTO USER(id,email,password,city,username,admin) VALUES(911113,"meeenz@meeenz.com","$argon2id$v=19$m=2048,t=1,p=2$m0Ro6ArcaMfanzBFGVmQCw$vmDrLnu2CfevEJwJh/KeVu53cScOfjYzF57jNIFPJ4Q","Mainz","meeenzman",FALSE);
--- Users

INSERT INTO USER_NOTIFICATIONS(user_id,info,date) VALUES(911111,"<p> GOODBYE WORLD </p>","2000-01-01");
-- Attraction
INSERT INTO ATTRACTION_ENTRY (id,title, type, recommended_count, city, info, approved, PosX, PosY, stars)
VALUES(9000,'Brandenburg Gate', 'Landmark', 12000, 'Berlin', 'An 18th-century neoclassical monument in Berlin, one of the best-known landmarks of Germany.', TRUE, 52.5163, 13.3777, 4.8),
(9001,'Neuschwanstein Castle', 'Castle', 15000, 'Schwangau', 'A 19th-century historicist palace on a rugged hill above the village of Hohenschwangau near Füssen in southwest Bavaria.', TRUE, 47.5576, 10.7498, 4.9),
(9002,'Cologne Cathedral', 'Cathedral', 18000, 'Cologne', 'A renowned monument of German Catholicism and Gothic architecture, also a World Heritage Site.', FALSE, 50.9413, 6.9583, 4.8),
(9003,'Heidelberg Castle', 'Castle', 8000, 'Heidelberg', 'A famous ruin in Germany and landmark of Heidelberg, built on a hill overlooking the city.', TRUE, 49.4106, 8.7158, 4.7),
(9004,'Miniatur Wunderland', 'Museum', 9000, 'Hamburg', 'The world\'s largest model railway exhibition, located in Hamburg.', TRUE, 53.5430, 9.9881, 4.9),
(9005,'Zugspitze', 'Mountain', 6000, 'Garmisch-Partenkirchen', 'The highest peak of the Wetterstein Mountains as well as the highest mountain in Germany.', TRUE, 47.4211, 10.9853, 4.8),
(9006,'Marienplatz', 'Square', 11000, 'Munich', 'A central square in the city center of Munich, featuring the New Town Hall and the Glockenspiel.', TRUE, 48.1374, 11.5755, 4.7),
(9007,'Sanssouci Palace', 'Palace', 7000, 'Potsdam', 'The former summer palace of Frederick the Great, King of Prussia, in Potsdam, near Berlin.', TRUE, 52.4011, 13.0416, 4.7),
(9008,'AMONGUS Palace', 'Palace', 7000, 'Imposter', 'The former summer palace of Frederick the SUS, King of SUSSEX, in SUSSEX.', FALSE, 12.4011, 13.0416, 0);
-- Attraction

-- Review
INSERT INTO ATTRACTION_REVIEW(user_id,attraction_id,text,stars) VALUES(911113,9000,"Great place for my trad wife and 50 kids would go again, was able to drink my beer in peace without my bitch wife nagging",5);


-- Review
-- TESTDATA