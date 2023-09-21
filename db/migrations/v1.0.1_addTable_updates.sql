\c

CREATE TABLE updates (
    id SERIAL PRIMARY KEY,
    job_id integer REFERENCES jobs,
    display VARCHAR NOT NULL,
    date_updated TIMESTAMP NOT NULL,
    notes VARCHAR NULL
);