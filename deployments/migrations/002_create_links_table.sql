CREATE TABLE IF NOT EXISTS links (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    short_url varchar(128) UNIQUE NOT NULL,
    origin_url varchar(128) UNIQUE NOT NULL
);
CREATE INDEX idx_short_url on links(short_url);
CREATE INDEX idx_original_url on links(origin_url);