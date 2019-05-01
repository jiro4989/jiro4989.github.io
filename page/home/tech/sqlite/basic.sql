DROP TABLE IF EXISTS user;
CREATE TABLE user(id INTEGER, name TEXT);
INSERT INTO user (id, name) VALUES (1, 'taro');
INSERT INTO user (id, name) VALUES (2, 'ichiro');
INSERT INTO user (id, name) VALUES (3, 'hanako');
INSERT INTO user (id, name) VALUES (4, 'hiroshi');

DROP TABLE IF EXISTS hobby;
CREATE TABLE hobby(id INTEGER, user_id, name TEXT, description TEXT);
INSERT INTO hobby (id, user_id, name, description) VALUES (1, 1, 'baseball', NULL);
INSERT INTO hobby (id, user_id, name, description) VALUES (2, 1, 'pingpong', NULL);
INSERT INTO hobby (id, user_id, name, description) VALUES (3, 2, 'game', NULL);
INSERT INTO hobby (id, user_id, name, description) VALUES (3, 2, 'dart', NULL);
INSERT INTO hobby (id, user_id, name, description) VALUES (4, 3, 'cooking', NULL);
INSERT INTO hobby (id, user_id, name, description) VALUES (5, 3, 'dance', NULL);
INSERT INTO hobby (id, user_id, name, description) VALUES (6, 4, 'music', NULL);
INSERT INTO hobby (id, user_id, name, description) VALUES (7, 4, 'volunteer', NULL);
