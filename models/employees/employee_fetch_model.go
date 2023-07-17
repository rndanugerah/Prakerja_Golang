package employees

import (
	"net/http"
	"prakerja/db"
	"prakerja/models"
)

func FetchAllEmployees() (models.Response, error) {
	var obj Employee
	var arrobj []Employee
	var res models.Response

	con := db.CreateCon()
 
	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	} 

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Email, &obj.Telepon)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
