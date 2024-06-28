CREATE TABLE drinks (
    id INT NOT NULL AUTO_INCREMENT,
    name_kr VARCHAR(255),
    name_en VARCHAR(500),
    img_url VARCHAR(500),
    kcal VARCHAR(100),
    sat_fat VARCHAR(100),
    protein VARCHAR(100),
    fat VARCHAR(100),
    trans_fat VARCHAR(100),
    sodium VARCHAR(100),
    sugars VARCHAR(100),
    caffeine VARCHAR(100),
    cholesterol VARCHAR(100),
    chabo VARCHAR(100),
    PRIMARY KEY (id)
)