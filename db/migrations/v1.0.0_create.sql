CREATE DATABASE jobs;

\c jobs

CREATE TABLE jobs (
  id SERIAL PRIMARY KEY,
  company VARCHAR NOT NULL,
  job_title VARCHAR NOT NULL,
  date_applied TIMESTAMP NOT NULL,
  job_url VARCHAR NULL
);