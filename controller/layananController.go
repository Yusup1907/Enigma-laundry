package controller

import (
	"database/sql"
	"enigma_laundry/config"
	"enigma_laundry/model"
	"fmt"
)

func CreateLayanan(layanan model.Layanan, tx *sql.Tx) error {
	if layanan.Nama_Layanan == "" || layanan.Harga == 0 || layanan.Satuan == "" {
		return fmt.Errorf("Field tidak boleh kosong")
	}
	if layanan.Harga < 0 {
		return fmt.Errorf("Harga layanan tidak boleh negatif")
	}

	queryInsert := "INSERT INTO mst_layanan (id_layanan, nama_layanan, harga, satuan) VALUES ($1, $2, $3, $4);"
	_, err := tx.Exec(queryInsert, layanan.Id_Layanan, layanan.Nama_Layanan, layanan.Harga, layanan.Satuan)
	if err != nil {
		return err
	}

	fmt.Println("Successfully inserted data")
	return nil
}

func GetAllLayanan() []model.Layanan {
	db := config.ConnectDb()
	defer db.Close()

	queryLayanan := "SELECT * FROM mst_layanan;"

	rows, err := db.Query(queryLayanan)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	layanans := []model.Layanan{}
	for rows.Next() {
		layanan := model.Layanan{}
		err := rows.Scan(&layanan.Id_Layanan, &layanan.Nama_Layanan, &layanan.Harga, &layanan.Satuan)
		if err != nil {
			panic(err)
		}
		layanans = append(layanans, layanan)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return layanans
}

func GetLayananById(id string, tx *sql.Tx) (*model.Layanan, error) {

	queryLayananById := "SELECT id_layanan, nama_layanan, harga, satuan FROM mst_layanan WHERE id_layanan = $1;"
	row := tx.QueryRow(queryLayananById, id)

	layanan := &model.Layanan{}
	err := row.Scan(&layanan.Id_Layanan, &layanan.Nama_Layanan, &layanan.Harga, &layanan.Satuan)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Layanan with ID %s not found", id)
		}
		return nil, err
	}

	return layanan, nil

}

func UpdateLayanan(layanan model.Layanan, tx *sql.Tx) error {
	if layanan.Nama_Layanan == "" || layanan.Harga == 0 || layanan.Satuan == "" {
		return fmt.Errorf("Field tidak boleh kosong")
	}

	if layanan.Harga < 0 {
		return fmt.Errorf("Harga layanan tidak boleh negatif")
	}

	updateLayanan := "UPDATE mst_layanan SET nama_layanan = $2, harga = $3, satuan = $4 WHERE id_layanan = $1;"

	var err error
	_, err = tx.Exec(updateLayanan, layanan.Id_Layanan, layanan.Nama_Layanan, layanan.Harga, layanan.Satuan)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successfully Update Data")
	}

	return nil

}

func DeleteLayanan(id string, tx *sql.Tx) error {
	exists, err := isLayananExists(id, tx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Layanan dengan ID %s tidak ditemukan", id)
	}

	hasTransactions, err := hasCustomerTransactions(id, tx)
	if err != nil {
		return err
	}
	if hasTransactions {
		return fmt.Errorf("Layanan dengan ID %s memiliki transaksi terkait dan tidak dapat dihapus", id)
	}

	query := "DELETE FROM mst_layanan WHERE id_layanan = $1;"
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	} else {
		fmt.Println("successfully Deleted Data")
	}

	return nil
}

// Fungsi untuk memeriksa keberadaan layanan berdasarkan ID
func isLayananExists(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM mst_layanan WHERE id_layanan = $1"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Fungsi untuk memeriksa apakah pelanggan memiliki transaksi yang terkait
func hasLayananTransactions(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM trx_order_detail WHERE layanan_id = $1"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
