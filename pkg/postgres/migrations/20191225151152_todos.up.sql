CREATE TABLE todos (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    note TEXT NOT NULL,
    due_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
