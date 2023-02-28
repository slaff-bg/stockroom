-- = = = = = = =
-- When Docker Container is used, there is a delay in installing the extensions.
-- Because of it, the seeds fail to fulfill. That's why I impose this check. 
-- = = = = = = =
DO $$
DECLARE
    extItem text;
    extItems text[] := ARRAY['uuid-ossp', 'pgcrypto']; -- Required extensions.
    extFound varchar(16);
    maxAttempts integer := 4; -- The maximum number of check attempts.
    attempts integer := 0;
    sleepSeconds integer := 2;
BEGIN
    FOREACH extItem IN ARRAY extItems LOOP
        -- RAISE INFO E'Start searching extension [%].', extItem;

        WHILE attempts < maxAttempts
        LOOP
            -- RAISE INFO 'Attempts: %', attempts;
            SELECT extname INTO extFound FROM pg_extension WHERE extname = extItem;
            IF extFound IS NOT NULL AND extFound = extItem
            THEN
                -- RAISE INFO E'The extension [%] was found.\n', extItem;
                attempts := maxAttempts;
            ELSE
                RAISE INFO 'The extension [%] was not found.', extItem;
                attempts := attempts + 1;
                PERFORM pg_sleep(sleepSeconds);
            END IF;
        END LOOP;

         attempts := 0;
    END LOOP;

END $$;


-- = = = = = = =
-- Structure
-- = = = = = = =

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS customers (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    brand_name text NOT NULL,
    created_at timestamp with time zone DEFAULT current_timestamp,
    updated_at timestamp with time zone DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id uuid NOT NULL,
    email VARCHAR(64) NOT NULL,
    passwd TEXT NOT NULL,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    created_at timestamp with time zone DEFAULT current_timestamp,
    updated_at timestamp with time zone DEFAULT current_timestamp,
    
    -- CONSTRAINT users_customer_id_email_unique UNIQUE (customer_id, email),
    CONSTRAINT users_customer_fkey FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX idx__users_customer_id ON users(customer_id, email);

-- Handles updating the timestamp of the records when changes occur.
CREATE OR REPLACE FUNCTION update_updated_at_column() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW; 
END;
$$ language 'plpgsql';

CREATE TRIGGER update_customers__updated_at BEFORE UPDATE ON customers FOR EACH ROW EXECUTE PROCEDURE  update_updated_at_column();
CREATE TRIGGER update_users__updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE  update_updated_at_column();


-- = = = = = = =
-- Seeds
-- = = = = = = =

INSERT INTO "customers"("brand_name")
VALUES
('Brand 1'),
('Brand 2')
;

INSERT INTO users("customer_id", "email", "passwd", "first_name", "last_name")
select
    unnest(ARRAY[c.id, c.id]),
    CONCAT(unnest(ARRAY['bruce-lee@', 'chackie-chan@']), LOWER(REPLACE(c.brand_name, ' ', '')), '.local'),
    crypt('password', gen_salt('bf')),
    unnest(ARRAY['Bruce', 'Jackie']) as first_name,
    unnest(ARRAY['Lee', 'Chan'])
from customers as c order by c.id asc
