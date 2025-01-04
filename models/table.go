package models

import "exceltodatabases/config"

type Table struct {
	Column1   string
	Column2   string
	Column3   string
	Column4   string
	Column5   string
	Column6   string
	Id_table1 int
}

type TableRespon struct {
	Id_table1 int
	Column1   string
	Column2   string
	Column3   string
	Column4   string
	Column5   string
	Column6   string
	Id_table2 int
}

func GetAll() []TableRespon {
	rows, err := config.DB.Query(`SELECT
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
	`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var tabel []TableRespon
	for rows.Next() {
		var product TableRespon
		if err := rows.Scan(
			&product.Id_table1,
			&product.Column1,
			&product.Column2,
			&product.Column3,
			&product.Id_table2,
			&product.Column4,
			&product.Column5,
			&product.Column6,
		); err != nil {
			panic(err)
		}
		tabel = append(tabel, product)
	}
	return tabel
}
