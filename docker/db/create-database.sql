CREATE DATABASE planner;
\c planner;
BEGIN;
    CREATE TABLE IF NOT EXISTS shift_details (
        id SERIAL PRIMARY KEY,
        from_time INTEGER,
        to_time INTEGER,
        status varchar(100)
    );

    CREATE TABLE IF NOT EXISTS worker_details (
        id SERIAL PRIMARY KEY,
        name varchar(200),
        email varchar(200) UNIQUE,
        phone INTEGER,
        status varchar(100),
    );

    CREATE TABLE IF NOT EXISTS worker_shift_details (
        id SERIAL PRIMARY KEY,
        worker_id INTEGER,
        shift_id INTEGER,
        date TIMESTAMP,
        status varchar(100),
        CONSTRAINT fk_worker
            FOREIGN KEY(worker_id) REFERENCES worker_details(id) ON DELETE CASCADE,
        CONSTRAINT fk_shift
            FOREIGN KEY(shift_id) REFERENCES shift_details(id) ON DELETE CASCADE
    );
COMMIT;
