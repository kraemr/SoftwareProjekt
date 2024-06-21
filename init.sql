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


INSERT INTO ATTRACTION_ENTRY (id, title, type, recommended_count, city, street, housenumber, info, approved, PosX, PosY, stars, img_url, added_by)
VALUES
(7, 'Brandenburger Tor', 'Historical', 1500, 'Berlin', 'Pariser Platz', '1', 'A neoclassical monument that has stood through the city\'s history since the 18th century.', TRUE, 52.516275, 13.377704, 5, 'https://lh5.googleusercontent.com/p/AF1QipNaifG9JhlSPzLGHOn6hFKSlGWaXXhaIrPeCMdU=w408-h272-k-no', 911111),
(8, 'Reichstag Building', 'Government', 1400, 'Berlin', 'Platz der Republik', '1', 'The meeting place of the German parliament, featuring a glass dome with a 360-degree view of the city.', TRUE, 52.518623, 13.376198, 5, 'https://lh5.googleusercontent.com/p/AF1QipMslBMDL6y2w_WQHAAinSht9ldHy2iIuIpIxWBU=w408-h288-k-no', 911111),
(9, 'Berlin Cathedral', 'Religious', 1200, 'Berlin', 'Am Lustgarten', '1', 'An iconic church known for its stunning architecture and a viewing gallery.', TRUE, 52.51937, 13.4014, 4.5, 'https://lh5.googleusercontent.com/p/AF1QipNe9oHmRT_9yC0dyFWoUxpJoCiBCFdfSoMeyukh=w408-h306-k-no', 911111),
(10, 'Museum Island', 'Museum', 1100, 'Berlin', 'Bodestraße', '', 'A complex of five internationally significant museums, listed as a UNESCO World Heritage Site.', TRUE, 52.5212, 13.3956, 4.5, 'https://lh3.googleusercontent.com/gps-proxy/ALd4DhG1iElcv9nNSkimQh7-wXWSd3ntdBjlfpFoUqxPEilIOHOvH9IMhHLQ_pUS4Ytns_Hna7SNjYqO1kaS6_yKv770NYrEtEfwS-p4NqdQscz4a9aOw7ji6nVCCYq3TK4AhmvqU1VyPb0xJtt83hORDTe74jSAg3-f-kwf5gLFRPisZEGYCqHSiJkr=w408-h304-k-no', 911111),
(11, 'Berlin Wall Memorial', 'Historical', 1000, 'Berlin', 'Bernauer Straße', '111', 'A central memorial site of German division, located in the middle of the capital.', TRUE, 52.5351, 13.3903, 4.5, 'https://lh5.googleusercontent.com/p/AF1QipNSFDClh2TNTnKWdW1r2PE0L03Lag30RseMUOM=w408-h260-k-no', 911111),
(12, 'Checkpoint Charlie', 'Historical', 950, 'Berlin', 'Friedrichstraße', '43-45', 'The best-known Berlin Wall crossing point between East and West Berlin during the Cold War.', TRUE, 52.5076, 13.3904, 4, 'https://lh5.googleusercontent.com/p/AF1QipPi--N1WM6OnIWEPiw3Tj-RYbSgf0utoFct-HfL=w408-h305-k-no', 911111),
(13, 'Alexanderplatz', 'Public Square', 900, 'Berlin', 'Alexanderplatz', '', 'A large public square and transport hub, named after Tsar Alexander I.', TRUE, 52.5219, 13.4132, 4, 'https://lh5.googleusercontent.com/p/AF1QipOrRqUC2R_uJO47jwX3eeyzz1jNYIZ4SqN9fkMA=w408-h544-k-no', 911111),
(14, 'East Side Gallery', 'Art', 850, 'Berlin', 'Mühlenstraße', '3-100', 'An open-air gallery on the longest surviving section of the Berlin Wall, covered in murals.', TRUE, 52.505, 13.439, 4.5, 'https://lh5.googleusercontent.com/p/AF1QipNuw9A1a9DqKAulZbQ-2iKqdrA5fvoNI6jNnVID=w408-h306-k-no', 911111),
(15, 'Potsdamer Platz', 'Square', 800, 'Berlin', 'Potsdamer Platz', '', 'An important public square and traffic intersection in the centre of Berlin.', TRUE, 52.5096, 13.3759, 4, 'https://lh5.googleusercontent.com/p/AF1QipMd5-rZShrua-Jz5V_SAqoueUx52Fe9LeiWIVJe=w408-h306-k-no', 911111),
(16, 'Charlottenburg Palace', 'Historical', 750, 'Berlin', 'Spandauer Damm', '20-24', 'The largest palace in Berlin, offering a glimpse into the baroque and rococo eras.', TRUE, 52.520, 13.295, 4.5, 'https://lh5.googleusercontent.com/p/AF1QipMbwQI4Jx9PSu1Gs6jb9ACuzta4Oc7aVw0S1jmC=w408-h306-k-no', 911111);



INSERT INTO ATTRACTION_ENTRY (id, title, type, recommended_count, city, street, housenumber, info, approved, PosX, PosY, stars, img_url, added_by)
VALUES
(17, 'Pergamon Museum', 'Museum', 700, 'Berlin', 'Bodestraße', '1-3', 'One of the most visited museums in Germany, featuring monumental buildings like the Pergamon Altar.', TRUE, 52.5214, 13.3965, 5, 'https://lh5.googleusercontent.com/p/AF1QipMcib9mI5NNy_eBH2yoQQjkr-f1pokfPNRBKYRG=w408-h271-k-no', 911111),
(18, 'Berlin Zoological Garden', 'Zoo', 650, 'Berlin', 'Hardenbergplatz', '8', 'The oldest and most famous zoo in Germany, home to a wide variety of species.', TRUE, 52.5075, 13.3372, 4.5, 'https://lh3.googleusercontent.com/gps-proxy/ALd4DhHfMFwGMyiMylWs_B2axBgIB3tQkuw1wvPvC7rZASmTKj9teIMcxKiPB7Bfex1Ua64Y0lXx8XbEW1JN5YAmdhlDJa2NN8wEQpR0xb5UztltIoXGiDnOb6PB9SmTvFjnjhUN3qbR88-c0KqIRgbNAyzRsfpjvaJ40KGCTd4M2wq-Uj5PLyT1K8Lk=w408-h272-k-no', 911111),
(19, 'Gendarmenmarkt', 'Square', 600, 'Berlin', 'Gendarmenmarkt', '', 'A picturesque square featuring the Konzerthaus, the French Cathedral, and the German Cathedral.', TRUE, 52.5139, 13.3924, 4.5, 'https://lh5.googleusercontent.com/p/AF1QipM1Xo0cruKQ8VIIsYj7CJcgFtzh7ST0j9N85RDC=w408-h724-k-no', 911111),
(20, 'Kurfürstendamm', 'Street', 550, 'Berlin', 'Kurfürstendamm', '', 'A famous avenue known for its shops, houses, and hotels, often considered the Champs-Élysées of Berlin.', TRUE, 52.5026, 13.3301, 4, 'https://lh5.googleusercontent.com/p/AF1QipPZl-Hy7CsXOKSN1VrKdhsGxC_sW1xnGm_nQutK=w408-h544-k-no', 911111),
(21, 'Victory Column', 'Monument', 500, 'Berlin', 'Großer Stern', '1', 'A monument commemorating the Prussian victory in the Danish-Prussian War, offering a panoramic view of Berlin.', TRUE, 52.5145, 13.3501, 4.5, 'https://lh5.googleusercontent.com/p/AF1QipMnPN3c-e81mpPLTwrXxpAiITzmfK64k4GiCp6_=w408-h544-k-no', 911111),
(22, 'Topography of Terror', 'Museum', 450, 'Berlin', 'Niederkirchnerstraße', '8', 'An indoor and outdoor museum documenting the terror of the Nazi regime.', TRUE, 52.5063, 13.3849, 4.5, 'https://lh5.googleusercontent.com/p/AF1QipPZ6UT85ELKmYOcfU59m6L_S_MyMexdY3U_vxx4=w408-h306-k-no', 911111),
(23, 'Berlin TV Tower', 'Observation', 400, 'Berlin', 'Panoramastraße', '1A', 'The tallest structure in Germany, offering an observation deck with a view of Berlin.', TRUE, 52.5208, 13.4094, 4.5, 'https://lh5.googleusercontent.com/p/AF1QipMASq3OP3DrDMYRhZteXS2_Qfd6m_q8rRrqHiPH=w408-h725-k-no', 911111),
(24, 'Bode Museum', 'Museum', 350, 'Berlin', 'Bodestraße', '1-3', 'Part of the Museum Island complex, featuring collections of sculptures, coins, and Byzantine art.', TRUE, 52.5225, 13.3953, 4, 'https://lh5.googleusercontent.com/p/AF1QipPE7vtnT3z8Ks0DpV_xlnQ7SXETF_Q8LdCa8TpW=w408-h306-k-no', 911111),
(25, 'Hackesche Höfe', 'Courtyard', 300, 'Berlin', 'Rosenthaler Straße', '40-41', 'A complex of interlinked courtyards in the Spandau district, known for its vibrant cultural scene.', TRUE, 52.5252, 13.4018, 4, 'https://lh5.googleusercontent.com/p/AF1QipMV-0MTPqcHHPqEfQtutmvYTPJdj2Cl9mVd4A-H=w408-h306-k-no', 911111),
(26, 'Neue Nationalgalerie', 'Museum', 250, 'Berlin', 'Potsdamer Straße', '50', 'A museum for modern art, showcasing works from the early 20th century.', TRUE, 52.5071, 13.3654, 4.5, 'https://lh5.googleusercontent.com/p/AF1QipN9_q2kyXQp_GnPlCi66tf4zlvxmfQ3JrAigsuz=w408-h408-k-no', 911111);



INSERT INTO ATTRACTION_ENTRY (id, title, type, recommended_count, city, street, housenumber, info, approved, PosX, PosY, stars, img_url, added_by)
VALUES
(27, 'Jewish Museum Berlin', 'Museum', 600, 'Berlin', 'Lindenstraße', '9-14', 'A museum covering two millennia of German-Jewish history.', TRUE, 52.5024, 13.3957, 4.5, 'https://lh5.googleusercontent.com/p/AF1QipN9_q2kyXQp_GnPlCi66tf4zlvxmfQ3JrAigsuz=w408-h408-k-no', 911111),
(28, 'Tiergarten', 'Park', 580, 'Berlin', 'Straße des 17. Juni', '', 'Berlin’s most popular inner-city park.', TRUE, 52.5146, 13.3501, 4.5, 'https://example.com/images/tiergarten.jpg', 911111),
(29, 'Charlottenburg Palace Gardens', 'Park', 550, 'Berlin', 'Spandauer Damm', '20-24', 'Beautiful baroque gardens surrounding the Charlottenburg Palace.', TRUE, 52.5201, 13.2951, 4.5, 'https://example.com/images/charlottenburg_palace_gardens.jpg', 911111),
(30, 'German Historical Museum', 'Museum', 500, 'Berlin', 'Unter den Linden', '2', 'A museum dedicated to German history, from its beginnings to the present.', TRUE, 52.5174, 13.3963, 4.5, 'https://example.com/images/german_historical_museum.jpg', 911111),
(31, 'Berlin Dungeon', 'Attraction', 480, 'Berlin', 'Spandauer Straße', '2', 'An interactive experience that brings Berlin’s dark history to life.', TRUE, 52.5186, 13.4054, 4, 'https://example.com/images/berlin_dungeon.jpg', 911111),
(32, 'Kaiser Wilhelm Memorial Church', 'Historical', 460, 'Berlin', 'Breitscheidplatz', '', 'A memorial church to commemorate Emperor Wilhelm I.', TRUE, 52.5076, 13.3368, 4.5, 'https://example.com/images/kaiser_wilhelm_memorial_church.jpg', 911111),
(33, 'Treptower Park', 'Park', 450, 'Berlin', 'Alt-Treptow', '', 'A large park in the Treptow-Köpenick district.', TRUE, 52.4877, 13.4716, 4.5, 'https://example.com/images/treptower_park.jpg', 911111),
(34, 'Sony Center', 'Complex', 430, 'Berlin', 'Potsdamer Platz', '2', 'A complex of buildings located at the Potsdamer Platz.', TRUE, 52.5096, 13.3731, 4, 'https://example.com/images/sony_center.jpg', 911111),
(35, 'Tempelhofer Feld', 'Park', 420, 'Berlin', 'Tempelhofer Damm', '', 'A park created on the site of the former Tempelhof Airport.', TRUE, 52.4751, 13.4025, 4.5, 'https://example.com/images/tempelhofer_feld.jpg', 911111),
(36, 'Mauerpark', 'Park', 410, 'Berlin', 'Bernauer Straße', '', 'A public park in the Prenzlauer Berg district, known for its flea market and karaoke sessions.', TRUE, 52.5437, 13.4023, 4.5, 'https://example.com/images/mauerpark.jpg', 911111);



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
