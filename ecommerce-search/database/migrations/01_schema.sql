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