CREATE TABLE users (
    uid UUID NOT NULL DEFAULT pg_catalog.gen_random_uuid(),
    username text NULL,
    email text NOT NULL,
    "password" text NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL,
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_pkey PRIMARY KEY (uid)
);
CREATE INDEX idx_users_deleted_at ON users USING btree (deleted_at);