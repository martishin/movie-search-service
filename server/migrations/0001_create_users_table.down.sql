-- Drop the trigger first (otherwise, the table drop will fail)
DROP TRIGGER IF EXISTS set_timestamp_users ON users;

-- Drop the users table
DROP TABLE IF EXISTS users;

-- Drop the function only if no other table uses it
DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT
                1
            FROM
                pg_trigger
            WHERE
                tgname = 'set_timestamp'
        ) THEN
            DROP FUNCTION IF EXISTS update_updated_at_column;
        END IF;
    END
$$;
