-- 02_schema.sql

-- Enable necessary extensions
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE categories (
  category_id SERIAL PRIMARY KEY,
  category_name VARCHAR(255) NOT NULL,
  parent_category_id INTEGER REFERENCES categories(category_id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
  product_id SERIAL PRIMARY KEY,
  product_name VARCHAR(255) NOT NULL,
  description TEXT,
  price DECIMAL(10,2) NOT NULL,
  stock_quantity INTEGER NOT NULL,
  search_vector TSVECTOR,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_categories (
  product_id INTEGER REFERENCES products(product_id),
  category_id INTEGER REFERENCES categories(category_id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (product_id, category_id)
);

CREATE TABLE tags (
  tag_id SERIAL PRIMARY KEY,
  tag_name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE product_tags (
  product_id INTEGER REFERENCES products(product_id),
  tag_id INTEGER REFERENCES tags(tag_id),
  PRIMARY KEY (product_id, tag_id)
);

CREATE TABLE synonyms (
  word VARCHAR(255) PRIMARY KEY,
  synonyms TEXT[]
);

CREATE INDEX idx_products_search_vector ON products USING GIN (search_vector);

-- Create trigger function
CREATE OR REPLACE FUNCTION products_search_update() RETURNS trigger AS $$
DECLARE
    search_text_with_synonyms TEXT;
BEGIN
    -- Combine product name and description
    search_text_with_synonyms := NEW.product_name || ' ' || COALESCE(NEW.description, '');
    
    -- Add synonyms
    SELECT search_text_with_synonyms || ' ' || string_agg(s.synonym, ' ')
    INTO search_text_with_synonyms
    FROM unnest(string_to_array(lower(search_text_with_synonyms), ' ')) AS word
    LEFT JOIN synonyms ON synonyms.word = word
    LEFT JOIN unnest(synonyms.synonyms) AS s(synonym) ON true
    GROUP BY search_text_with_synonyms;

    NEW.search_vector := setweight(to_tsvector('english', NEW.product_name), 'A') ||
                         setweight(to_tsvector('english', COALESCE(NEW.description, '')), 'B') ||
                         setweight(to_tsvector('simple', NEW.product_name), 'A') ||
                         setweight(to_tsvector('simple', COALESCE(NEW.description, '')), 'B');
    NEW.search_text := search_text_with_synonyms;
    RETURN NEW;
END
$$ LANGUAGE plpgsql;


-- Create trigger
CREATE TRIGGER products_search_update
BEFORE INSERT OR UPDATE ON products
FOR EACH ROW EXECUTE FUNCTION products_search_update();

-- Initial update for existing products
UPDATE products 
SET updated_at = CURRENT_TIMESTAMP;