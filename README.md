# gojobsapi
An API to track job applications written in Go.

## Setup

1. Install Docker
2. run `docker compose up`
3. Install psql client
    - `sudo apt-get install postgresql-client-common`
    - `sudo apt-get install postgresql-client`
3. Create DB
    - Connect to DB
        - `psql -h localhost -U <db username> -d jobs`
    - Run migrations
        - `\i <path to migration file>`

The server runs on port 3000 by default
