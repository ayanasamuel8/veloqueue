CREATE TYPE job_state AS ENUM (
    'PENDING',
    'RUNNING',
    'FAILED',
    'DEAD',
    'COMPLETED'
);

CREATE TABLE jobs(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    queue TEXT NOT NULL DEFAULT 'default',
    kind TEXT NOT NULL,
    args JSONB,
    state job_state NOT NULL DEFAULT 'PENDING',
    scheduled_for TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    locked_until TIMESTAMPTZ,
    attempts INT NOT NULL DEFAULT 0,
    max_attempts INT NOT NULL DEFAULT 25,
    last_error TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    started_at TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);