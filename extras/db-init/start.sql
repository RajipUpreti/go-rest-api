-- create new database named kong --
CREATE DATABASE kong;

-- create new user named kong with password --
CREATE USER kong WITH PASSWORD 'kong';

-- grant all privileges to the kong user on the kong database --
GRANT ALL PRIVILEGES ON DATABASE kong TO kong;

-- grant all privileges to the kong user on all schemas, tables, and sequences in the kong database --
\c kong
GRANT ALL PRIVILEGES ON SCHEMA public TO kong;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO kong;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO kong;

-- ensure future tables and sequences also have privileges --
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO kong;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON SEQUENCES TO kong;
