CREATE TYPE type AS ENUM ('Major', 'Minor', '7', '5', 'dim');

CREATE TABLE chords (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  root text NOT NULL,
  type type NOT NULL,
  guitar json[],
  piano json[]
);