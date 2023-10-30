CREATE or replace TABLE users
(
	username VARCHAR(10) NOT NULL,
	`password` VARCHAR(10) NOT NULL,
	PRIMARY KEY (username)
);

CREATE OR REPLACE TABLE barangs
(
	ID INT AUTO_INCREMENT,
	nama VARCHAR(50),
	harga INT,
	stok_barang INT,
	pegawai VARCHAR(50),
	PRIMARY KEY (ID),
	FOREIGN KEY (pegawai) REFERENCES users(username)
);

CREATE OR REPLACE	TABLE customers
(
	ID INT AUTO_INCREMENT,
	nama VARCHAR(50),
	alamat_customer VARCHAR(50),
	PRIMARY KEY (ID)
);

CREATE OR REPLACE	TABLE customers_backup
(
	ID INT AUTO_INCREMENT,
	nama VARCHAR(50),
	alamat_customer VARCHAR(50),
	PRIMARY KEY (ID)
);

CREATE OR REPLACE TABLE nota_transactions
(
	ID INT AUTO_INCREMENT,
	username VARCHAR(50),
	barang_id INT,
	qty INT,
	customer_id INT,
	created_at DATE DEFAULT CURRENT_DATE(),
	PRIMARY KEY (ID),
	FOREIGN KEY (username) REFERENCES users(username),
	FOREIGN KEY (barang_id) REFERENCES barangs(ID),
	FOREIGN KEY (customer_id) REFERENCES customers(ID)
);

CREATE OR REPLACE TABLE nota_transactions_backup
(
	ID INT AUTO_INCREMENT,
	username VARCHAR(50),
	barang_id INT,
	qty INT,
	customer_id INT,
	created_at DATE DEFAULT CURRENT_DATE(),
	PRIMARY KEY (ID),
	FOREIGN KEY (username) REFERENCES users(username),
	FOREIGN KEY (barang_id) REFERENCES barangs(ID),
	FOREIGN KEY (customer_id) REFERENCES customers(ID)
);

INSERT INTO users
(username, `password`)
VALUES ('admin', 'admin')
;

select created_at(date,'%y-%m-%d') from nota_transactions;