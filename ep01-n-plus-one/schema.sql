-- A tiny shop: users write reviews on products.
-- NOTE: there is deliberately NO index on reviews.product_id yet.
-- Postgres does not create indexes for foreign keys automatically. Step 4 adds it.

CREATE TABLE users (
  id   serial PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE products (
  id           serial PRIMARY KEY,
  name         text NOT NULL,
  price_cents  int  NOT NULL,
  review_count int  NOT NULL DEFAULT 0   -- used in Step 7 (denormalized counter)
);

CREATE TABLE reviews (
  id         serial PRIMARY KEY,
  product_id int NOT NULL REFERENCES products(id),
  user_id    int NOT NULL REFERENCES users(id),
  rating     int NOT NULL,
  body       text NOT NULL,
  created_at timestamptz NOT NULL
);

SELECT setseed(0.42); -- same random data on every machine

INSERT INTO users (name)
SELECT 'user ' || g FROM generate_series(1, 500) g;

INSERT INTO products (name, price_cents)
SELECT 'product ' || g, (random() * 10000)::int FROM generate_series(1, 1000) g;

-- ~50 reviews per product on average, 50,000 total
INSERT INTO reviews (product_id, user_id, rating, body, created_at)
SELECT (random() * 999)::int + 1,
       (random() * 499)::int + 1,
       (random() * 4)::int + 1,
       'review text number ' || g,
       now() - random() * interval '365 days'
FROM generate_series(1, 50000) g;

-- Needed for the "Detect it" step
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;

ANALYZE;
