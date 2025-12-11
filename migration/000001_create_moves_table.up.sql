CREATE TABLE moves (
	id INT IDENTITY(1,1) PRIMARY KEY,
	name NVARCHAR(200) NOT NULL,
	type INT NOT NULL,
	category INT NOT NULL,
	power INT NULL,
	accuracy INT NULL,
	description NVARCHAR(MAX) NULL
);

