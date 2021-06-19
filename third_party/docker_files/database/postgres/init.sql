CREATE DATABASE ocp_note_api;
CREATE USER best_user WITH PASSWORD 'best_password';
GRANT all privileges ON DATABASE ocp_note_api TO best_user;

\c ocp_note_api best_user

CREATE TABLE notes (
    id SERIAL PRIMARY KEY, 
    user_id INT NOT NULL, 
    classroom_id INT NOT NULL, 
    document_id INT NOT NULL
);

ALTER TABLE notes OWNER TO best_user;
