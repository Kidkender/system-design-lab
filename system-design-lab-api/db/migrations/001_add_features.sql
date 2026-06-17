-- Migration: add new feature columns
-- Run once against existing database

ALTER TYPE session_status ADD VALUE IF NOT EXISTS 'abandoned';

DO $$ BEGIN
    CREATE TYPE session_mode AS ENUM ('normal', 'interview');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

ALTER TABLE steps
    ADD COLUMN IF NOT EXISTS hint TEXT;

ALTER TABLE scenarios
    ADD COLUMN IF NOT EXISTS time_limit_seconds INT;

ALTER TABLE user_sessions
    ADD COLUMN IF NOT EXISTS mode session_mode NOT NULL DEFAULT 'normal',
    ADD COLUMN IF NOT EXISTS completed_at TIMESTAMP;
