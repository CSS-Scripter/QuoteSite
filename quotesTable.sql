CREATE TABLE quotes(
  quote_id SERIAL,
  server_id VARCHAR(25),
  quote TEXT,
  by VARCHAR(255),
  year VARCHAR(5),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
