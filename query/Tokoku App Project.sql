CREATE or replace TABLE users
(
	username VARCHAR(10) NOT NULL,
	`password` VARCHAR(10) NOT NULL,
	PRIMARY KEY (username)
);

CREATE OR REPLACE TABLE barang
(
	ID INT AUTO_INCREMENT,
	nama VARCHAR(50),
	harga INT,
	stok_barang INT,
	PRIMARY KEY (ID)
);

CREATE OR REPLACE	TABLE customer
(
	ID INT AUTO_INCREMENT,
	nama VARCHAR(50),
	alamat_customer VARCHAR(50),
	PRIMARY KEY (ID)
);

CREATE OR REPLACE TABLE nota_transaksi
(
	ID INT AUTO_INCREMENT,
	username VARCHAR(50),
	barang_id INT,
	qty INT,
	customer_id INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
	PRIMARY KEY (ID),
	FOREIGN KEY (username) REFERENCES users(username),
	FOREIGN KEY (barang_id) REFERENCES barang(ID),
	FOREIGN KEY (customer_id) REFERENCES customer(ID)
);

INSERT INTO users
(username, `password`)
VALUES ('admin', 'admin')
;