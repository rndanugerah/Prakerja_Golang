package employees

import (
	"database/sql"
	"net/http"
	"prakerja/db"
	"prakerja/models"
)

func GetEmployeeByID(id int) (models.Response, error) {
	var obj Pegawai
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pegawai WHERE id = ?"

	row := con.QueryRow(sqlStatement, id)

	err := row.Scan(&obj.Id, &obj.Nama, &obj.Email, &obj.Telepon)
	if err != nil {
		if err == sql.ErrNoRows {
			// Jika data pegawai tidak ditemukan
			res.Status = http.StatusNotFound
			res.Message = "Employee not found"
			return res, nil
		}
		// Jika terjadi kesalahan lain saat memindai data
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}
