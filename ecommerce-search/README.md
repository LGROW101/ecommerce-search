# SELECT
```
SELECT * FROM categories;
SELECT * FROM products;
SELECT * FROM tags;
SELECT * FROM product_categories;
SELECT * FROM product_tags;
SELECT * FROM synonyms;


```
# INSERT
## ตาราง categories
```
INSERT INTO categories (category_name, parent_category_id) VALUES
('Electronics', NULL),
('Clothing', NULL),
('Books', NULL),
('Home & Kitchen', NULL),
('Sports & Outdoors', NULL),
('Smartphones', 1),
('Laptops', 1),
('T-shirts', 2),
('Jeans', 2),
('Fiction', 3);

```

##  ตาราง products
```
INSERT INTO products (product_name, description, price, stock_quantity) VALUES
('iPhone 12', 'Latest Apple smartphone', 799.99, 50),
('Samsung Galaxy S21', 'Flagship Android phone', 699.99, 40),
('Dell XPS 13', 'Powerful ultrabook', 1299.99, 30),
('Levi''s 501 Jeans', 'Classic straight fit jeans', 59.99, 100),
('Harry Potter Box Set', 'Complete series of Harry Potter books', 89.99, 25),
('Nike Air Max', 'Comfortable running shoes', 129.99, 60),
('Instant Pot Duo', 'Multi-functional pressure cooker', 99.99, 45),
('Canon EOS R', 'Full-frame mirrorless camera', 1799.99, 15),
('PlayStation 5', 'Next-gen gaming console', 499.99, 10),
('Kindle Paperwhite', 'E-reader with backlight', 129.99, 35);
```
##  ตาราง tags
```
INSERT INTO tags (tag_name) VALUES
('smartphone'),
('laptop'),
('denim'),
('book series'),
('running'),
('kitchen appliance'),
('camera'),
('gaming'),
('e-reader'),
('bestseller');

```

## ตาราง product_categories
```
INSERT INTO product_categories (product_id, category_id) VALUES
(1, 6), -- iPhone 12 in Smartphones
(2, 6), -- Samsung Galaxy S21 in Smartphones
(3, 7), -- Dell XPS 13 in Laptops
(4, 9), -- Levi's 501 Jeans in Jeans
(5, 10), -- Harry Potter Box Set in Fiction
(6, 5), -- Nike Air Max in Sports & Outdoors
(7, 4), -- Instant Pot Duo in Home & Kitchen
(8, 1), -- Canon EOS R in Electronics
(9, 1), -- PlayStation 5 in Electronics
(10, 3); -- Kindle Paperwhite in Books

```

##  ตาราง product_tags
```
INSERT INTO product_tags (product_id, tag_id) VALUES
(1, 1), -- iPhone 12 - smartphone
(2, 1), -- Samsung Galaxy S21 - smartphone
(3, 2), -- Dell XPS 13 - laptop
(4, 3), -- Levi's 501 Jeans - denim
(5, 4), -- Harry Potter Box Set - book series
(5, 10), -- Harry Potter Box Set - bestseller
(6, 5), -- Nike Air Max - running
(7, 6), -- Instant Pot Duo - kitchen appliance
(8, 7), -- Canon EOS R - camera
(9, 8), -- PlayStation 5 - gaming
(10, 9); -- Kindle Paperwhite - e-reader

```
## ตาราง synonyms
```
INSERT INTO synonyms (word, synonyms) VALUES
('smartphone', ARRAY['mobile', 'cellphone', 'handset']),
('laptop', ARRAY['notebook', 'portable computer']),
('jeans', ARRAY['denim', 'trousers']),
('book', ARRAY['novel', 'publication', 'tome']),
('shoes', ARRAY['footwear', 'sneakers']),
('cooker', ARRAY['stove', 'oven']),
('camera', ARRAY['photographic equipment', 'imaging device']),
('console', ARRAY['gaming system', 'video game platform']),
('e-reader', ARRAY['digital book reader', 'electronic book device']),
('bestseller', ARRAY['top-selling', 'popular book']);

```

# เชื่อมต่อกับ container ของ PostgreSQL

```

docker exec -it ecommerce-search-db-1 psql -U postgres -d ecommerce_search

```

# ทดสอบ api ด้วย curl 

```
curl "http://localhost:8080/search?q=smartphone"
curl "http://localhost:8080/search?q=laptop"
curl "http://localhost:8080/search?q=book"

```

# # การใช้ options เพิ่มเติมกับ curl

## แสดงส่วนหัว (headers) ของ response

```
curl -i "http://localhost:8080/search?q=smartphone"

```

## ใช้ verbose mode เพื่อดูรายละเอียดเพิ่มเติม

```
curl -v "http://localhost:8080/search?q=smartphone"

```



