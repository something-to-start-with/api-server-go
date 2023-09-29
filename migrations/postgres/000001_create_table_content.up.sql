CREATE TABLE IF NOT EXISTS content
(
    id   BIGSERIAL NOT NULL,
    data TEXT,
    CONSTRAINT pk_content PRIMARY KEY (id)
);