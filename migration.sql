CREATE DATABASE excel_to_databases use excel_to_databases 

CREATE Table table1(
    id INT PRIMARY KEY AUTO_INCREMENT,
    column1 VARCHAR(250),
    column2 VARCHAR(250),
    column3 VARCHAR(250)
)



CREATE TABLE table2 (
    id INT PRIMARY KEY AUTO_INCREMENT,
    column1 VARCHAR(250),
    column2 VARCHAR(250),
    column3 VARCHAR(250),
    id_table1 INT,
    FOREIGN KEY (id_table1) REFERENCES table1(id) ON DELETE CASCADE
);

drop Table table1
SELECT
    *
FROM
    table1
SELECT
    *
FROM
    table2


SELECT
    table1.id as id_table1,
    table1.column1,
    table1.column2,
    table1.column3,
    table2.id as id_table2,
    table2.column1 as column4,
    table2.column2 as column5,
    table2.column3 as column6
FROM
    table2
    INNER JOIN table1 ON table1.id = table2.id_table1



DELETE FROM table1

DROP TABLE table1

DROP TABLE table2
