CREATE TABLE IF NOT EXISTS rooms (
  "id" UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid (),
  "theme" VARCHAR NOT NULL
);

---- create above / drop below ----
DROP TABLE IF EXISTS rooms;
