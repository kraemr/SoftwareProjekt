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
CREATE INDEX _ATTRACTION_ENTRY_INDEX ON ATTRACTION_ENTRY(city);


-- To get Username JOIN
-- TODO Add DATE
CREATE TABLE IF NOT EXISTS ATTRACTION_REVIEW (
    id SERIAL PRIMARY KEY,   
    user_id INT,
    attraction_id INT,
    text TEXT,
    stars FLOAT
);
CREATE INDEX _ATTRACTION_REVIEW_INDEX ON ATTRACTION_REVIEW(attraction_id);




CREATE TABLE USER_NOTIFICATIONS(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    info TEXT NOT NULL,
    date DATE
);
CREATE INDEX _USER_NOTIFICATIONS_INDEX ON USER_NOTIFICATIONS(user_id);


-- add foreign key constraints
ALTER TABLE ATTRACTION_REVIEW ADD CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);
ALTER TABLE ATTRACTION_REVIEW ADD CONSTRAINT attraction_id_fk FOREIGN KEY (attraction_id) REFERENCES ATTRACTION_ENTRY(id);
ALTER TABLE USER_NOTIFICATIONS ADD CONSTRAINT notif_user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);
ALTER TABLE USER_PREFERENCES ADD CONSTRAINT pref_user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);

ALTER TABLE USER_FAVORITE ADD CONSTRAINT fav_user_id_fk FOREIGN KEY (user_id) REFERENCES USER(id);
ALTER TABLE USER_FAVORITE ADD CONSTRAINT fav_attraction_id_fk FOREIGN KEY (attraction_id) REFERENCES ATTRACTION_ENTRY(id);



-- TESTDATA
INSERT INTO USER(id,email,password,city,username,admin) VALUES(911111,"admin@testemail.com","$argon2id$v=19$m=2048,t=1,p=2$m0Ro6ArcaMfanzBFGVmQCw$vmDrLnu2CfevEJwJh/KeVu53cScOfjYzF57jNIFPJ4Q","Oppenheim","adminman",TRUE);
INSERT INTO USER(id,email,password,city,username,admin) VALUES(911112,"test@testemail.com","$argon2id$v=19$m=2048,t=1,p=2$m0Ro6ArcaMfanzBFGVmQCw$vmDrLnu2CfevEJwJh/KeVu53cScOfjYzF57jNIFPJ4Q","Müllhausen","testman",FALSE);
INSERT INTO USER(id,email,password,city,username,admin) VALUES(911113,"meeenz@meeenz.com","$argon2id$v=19$m=2048,t=1,p=2$m0Ro6ArcaMfanzBFGVmQCw$vmDrLnu2CfevEJwJh/KeVu53cScOfjYzF57jNIFPJ4Q","Mainz","meeenzman",FALSE);

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
(9007,'Sanssouci Palace', 'Palace', 7000, 'Potsdam', 'The former summer palace of Frederick the Great, King of Prussia, in Potsdam, near Berlin.', TRUE, 52.4011, 13.0416, 4.7);

INSERT INTO ATTRACTION_ENTRY (id, title, type, recommended_count, city, info, approved, PosX, PosY, stars)
VALUES
(9008, 'Berlin TV Tower', 'Observation Tower', 13000, 'Berlin', 'A television tower in central Berlin, offering panoramic views of the city.', TRUE, 52.5208, 13.4094, 4.6),
(9009, 'Museum Island', 'Museum Complex', 14000, 'Berlin', 'A complex of five internationally significant museums, part of the UNESCO World Heritage.', TRUE, 52.5169, 13.4010, 4.8),
(9010, 'Berlin Wall Memorial', 'Historical Site', 12000, 'Berlin', 'A memorial site of the Berlin Wall, featuring remnants of the wall and a documentation center.', TRUE, 52.5351, 13.3904, 4.7),
(9011, 'Charlottenburg Palace', 'Palace', 9000, 'Berlin', 'The largest palace in Berlin, a fine example of baroque and rococo architecture.', TRUE, 52.5204, 13.2952, 4.6),
(9012, 'Pergamon Museum', 'Museum', 11000, 'Berlin', 'One of the most visited museums in Berlin, housing monumental buildings such as the Pergamon Altar.', TRUE, 52.5213, 13.3969, 4.8),
(9013, 'Berlin Cathedral', 'Cathedral', 10000, 'Berlin', 'An iconic Protestant cathedral and one of the major landmarks on Museum Island.', TRUE, 52.5194, 13.4010, 4.7),
(9014, 'Checkpoint Charlie', 'Historical Site', 9500, 'Berlin', 'The best-known Berlin Wall crossing point between East and West Berlin during the Cold War.', TRUE, 52.5076, 13.3904, 4.5),
(9015, 'East Side Gallery', 'Art Gallery', 8000, 'Berlin', 'An open-air gallery on a long section of the Berlin Wall, featuring artworks from artists around the world.', TRUE, 52.5053, 13.4390, 4.6),
(9016, 'Potsdamer Platz', 'Square', 8500, 'Berlin', 'A bustling public square and traffic intersection in the center of Berlin.', TRUE, 52.5096, 13.3755, 4.5),
(9017, 'Gendarmenmarkt', 'Square', 7000, 'Berlin', 'A beautiful square in Berlin, featuring the Berlin Concert Hall and the French and German Churches.', TRUE, 52.5138, 13.3927, 4.7),
(9018, 'Kurfürstendamm', 'Shopping Street', 10000, 'Berlin', 'A famous avenue known for its luxury shops, hotels, and theaters.', TRUE, 52.5020, 13.3307, 4.6),
(9019, 'Berlin Zoological Garden', 'Zoo', 9000, 'Berlin', 'The oldest and most famous zoo in Germany, home to a wide variety of species.', TRUE, 52.5086, 13.3373, 4.7),
(9020, 'Victory Column', 'Monument', 8000, 'Berlin', 'A monument in Berlin commemorating the Prussian victory in the Danish-Prussian War.', TRUE, 52.5145, 13.3501, 4.6),
(9021, 'Berlin Dungeon', 'Attraction', 7500, 'Berlin', 'An immersive and interactive experience showcasing the dark history of Berlin.', TRUE, 52.5194, 13.4084, 4.4),
(9022, 'Tempelhofer Feld', 'Park', 7000, 'Berlin', 'A large public park and former airport in Berlin, popular for outdoor activities.', TRUE, 52.4750, 13.4050, 4.5),
(9023, 'Treptower Park', 'Park', 6500, 'Berlin', 'A large park alongside the river Spree, featuring the Soviet War Memorial.', TRUE, 52.4933, 13.4692, 4.5),
(9024, 'Hackescher Markt', 'Market', 6000, 'Berlin', 'A lively area known for its vibrant nightlife, shops, and restaurants.', TRUE, 52.5232, 13.4020, 4.4);

INSERT INTO USER_FAVORITE(user_id,attraction_id) VALUES(911111,9000); 
-- Review
INSERT INTO ATTRACTION_REVIEW(user_id,attraction_id,text,stars) VALUES(911113,9000,"Great place for my trad wife and 50 kids would go again, was able to drink my beer in peace without my bitch wife nagging",5);


-- Review
-- TESTDATA