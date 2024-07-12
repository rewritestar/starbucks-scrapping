CREATE TABLE drinks (
    id INT NOT NULL AUTO_INCREMENT,
    name_kr VARCHAR(255),
    name_en VARCHAR(500),
    img_url VARCHAR(500),
    kcal INT,
    sat_fat INT,
    protein INT,
    fat INT,
    trans_fat INT,
    sodium INT,
    sugars INT,
    caffeine INT,
    cholesterol INT,
    chabo INT,
    PRIMARY KEY (id)
)