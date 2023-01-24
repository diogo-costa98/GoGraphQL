--Execute sqlite3.exe homework.db ".read homework.sql" to obtain the desired sqlite database

CREATE TABLE users(     id          INTEGER PRIMARY KEY,
                        name        TEXT,
                        password    TEXT);

CREATE TABLE questions( id          INTEGER PRIMARY KEY,
                        body        TEXT,
                        user_id     INTEGER,
                        FOREIGN KEY(user_id) REFERENCES users(id));

CREATE TABLE options(   id          INTEGER PRIMARY KEY,
                        body        TEXT,
                        correct     INTEGER,
                        question_id INTEGER,
                        FOREIGN KEY(question_id) REFERENCES questions(id));

-- Insert data into the users table 
INSERT INTO users (name, password) VALUES ('user',          '$2y$14$wUAFYLyV86QWVdQZnkxXren38kLav0L1StZhcTXbTpAY/rNqav80e');                 -- Password is 'password'
INSERT INTO users (name, password) VALUES ('anotheruser',   '$2y$14$eBJ.BI.8ubr4jKhKZRhEUOuzRr.hI70kLCJOTGkcegqOGBpGvyxku');                 -- Password is 'anotherpassword'
INSERT INTO users (name, password) VALUES ('finaluser',     '$2y$14$rVvf9bcCf6sKKijwEz2FhOAh.Hj66W8hgVksTF5TRgekrZnNmrmuq');                 -- Password is 'finalpassword'

-- Insert data into the questions table
--user has 3 questions
INSERT INTO questions (body, user_id) VALUES ('How much is 1+1?', 1);
INSERT INTO questions (body, user_id) VALUES ('How much is 1+2?', 1);
INSERT INTO questions (body, user_id) VALUES ('How much is 1+3?', 1);
--anotheruser has 4 questions
INSERT INTO questions (body, user_id) VALUES ('How much is 2+1?', 2);
INSERT INTO questions (body, user_id) VALUES ('How much is 2+2?', 2);
INSERT INTO questions (body, user_id) VALUES ('How much is 2+3?', 2);
INSERT INTO questions (body, user_id) VALUES ('How much is 2+4?', 2);
--finaluser has 5 questions
INSERT INTO questions (body, user_id) VALUES ('How much is 3+1?', 3);
INSERT INTO questions (body, user_id) VALUES ('How much is 3+2?', 3);
INSERT INTO questions (body, user_id) VALUES ('How much is 3+3?', 3);
INSERT INTO questions (body, user_id) VALUES ('How much is 3+4?', 3);
INSERT INTO questions (body, user_id) VALUES ('How much is 3+5?', 3);

-- Insert data into the options table (each question has between 2 and 4 options
INSERT INTO options (body, correct, question_id) VALUES ('1+1=1', 0, 1);
INSERT INTO options (body, correct, question_id) VALUES ('1+1=2', 1, 1);
INSERT INTO options (body, correct, question_id) VALUES ('1+1=3', 0, 1);

INSERT INTO options (body, correct, question_id) VALUES ('1+2=1', 0, 2);
INSERT INTO options (body, correct, question_id) VALUES ('1+2=2', 0, 2);
INSERT INTO options (body, correct, question_id) VALUES ('1+2=3', 1, 2);
INSERT INTO options (body, correct, question_id) VALUES ('1+2=4', 0, 2);

INSERT INTO options (body, correct, question_id) VALUES ('1+3=3', 0, 3);
INSERT INTO options (body, correct, question_id) VALUES ('1+3=4', 1, 3);

INSERT INTO options (body, correct, question_id) VALUES ('2+1=1', 0, 4);
INSERT INTO options (body, correct, question_id) VALUES ('2+1=2', 0, 4);
INSERT INTO options (body, correct, question_id) VALUES ('2+1=3', 1, 4);

INSERT INTO options (body, correct, question_id) VALUES ('2+2=1', 0, 5);
INSERT INTO options (body, correct, question_id) VALUES ('2+2=2', 0, 5);
INSERT INTO options (body, correct, question_id) VALUES ('2+2=3', 0, 5);
INSERT INTO options (body, correct, question_id) VALUES ('2+2=4', 1, 5);

INSERT INTO options (body, correct, question_id) VALUES ('2+3=4', 0, 6);
INSERT INTO options (body, correct, question_id) VALUES ('2+3=5', 1, 6);

INSERT INTO options (body, correct, question_id) VALUES ('2+4=4', 0, 7);
INSERT INTO options (body, correct, question_id) VALUES ('2+4=5', 0, 7);
INSERT INTO options (body, correct, question_id) VALUES ('2+4=6', 1, 7);
INSERT INTO options (body, correct, question_id) VALUES ('2+4=7', 0, 7);

INSERT INTO options (body, correct, question_id) VALUES ('3+1=4', 1, 8);
INSERT INTO options (body, correct, question_id) VALUES ('3+1=5', 0, 8);

INSERT INTO options (body, correct, question_id) VALUES ('3+2=3', 0, 9);
INSERT INTO options (body, correct, question_id) VALUES ('3+2=2', 0, 9);
INSERT INTO options (body, correct, question_id) VALUES ('3+2=5', 1, 9);
INSERT INTO options (body, correct, question_id) VALUES ('3+2=4', 0, 9);

INSERT INTO options (body, correct, question_id) VALUES ('3+3=5', 0, 10);
INSERT INTO options (body, correct, question_id) VALUES ('3+3=6', 1, 10);

INSERT INTO options (body, correct, question_id) VALUES ('3+4=5', 0, 11);
INSERT INTO options (body, correct, question_id) VALUES ('3+4=6', 0, 11);
INSERT INTO options (body, correct, question_id) VALUES ('3+4=7', 1, 11);

INSERT INTO options (body, correct, question_id) VALUES ('3+5=8', 1, 12);
INSERT INTO options (body, correct, question_id) VALUES ('3+5=9', 0, 12);
