ALTER DATABASE root
  SET timezone TO 'Asia/Tokyo';

CREATE TABLE feeds (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  hyperlink TEXT NOT NULL UNIQUE,
  category INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL
);