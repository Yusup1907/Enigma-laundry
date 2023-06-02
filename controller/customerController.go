package controller

import (
	"database/sql"
	"enigma_laundry/config"
	"enigma_laundry/model"
	"fmt"
)

func CreateCustomer(customer model.Customer, tx *sql.Tx) error {
	if customer.Name == "" {
		return fmt.Errorf("Nama tidak boleh kosong")
	}

	if len(customer.Name) < 2 || len(customer.Name) > 40 {
		return fmt.Errorf("Nama harus terdiri dari 2 hingga 40 karakter")
	}

	if len(customer.No_Telp) < 10 || len(customer.No_Telp) > 12 {
		return fmt.Errorf("Nomor telepon harus terdiri dari 10 hingga 12 angka")
	}

	queryInsert := "INSERT INTO mst_customers (id_customer, name, no_telp, alamat) VALUES ($1, $2, $3, $4);"
	_, err := tx.Exec(queryInsert, customer.Id_Customer, customer.Name, customer.No_Telp, customer.Alamat)
	if err != nil {
		return err
	}

	fmt.Println("Successfully inserted data")
	return nil
}

func GetAllCustomer() []model.Customer {
	db := config.ConnectDb()
	defer db.Close()

	queryCustomer := "SELECT * FROM mst_customers;"

	rows, err := db.Query(queryCustomer)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := []model.Customer{}
	for rows.Next() {
		customer := model.Customer{}
		err := rows.Scan(&customer.Id_Customer, &customer.Name, &customer.No_Telp, &customer.Alamat)
		if err != nil {
			panic(err)
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return customers
}

func GetCustomerById(id string, tx *sql.Tx) (*model.Customer, error) {

	queryCustomerById := "SELECT id_customer, name, no_telp, alamat FROM mst_customers WHERE id_customer = $1;"
	row := tx.QueryRow(queryCustomerById, id)

	customer := &model.Customer{}
	err := row.Scan(&customer.Id_Customer, &customer.Name, &customer.No_Telp, &customer.Alamat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Customer with ID %s not found", id)
		}
		return nil, err
	}

	return customer, nil

}

func UpdateCustomer(customer model.Customer, tx *sql.Tx) error {
	if customer.Name == "" {
		return fmt.Errorf("Nama tidak boleh kosong")
	}

	if len(customer.Name) < 2 || len(customer.Name) > 40 {
		return fmt.Errorf("Nama harus terdiri dari 2 hingga 40 karakter")
	}

	updateCustomer := "UPDATE mst_customers SET name = $2, no_telp = $3, alamat = $4 WHERE id_customer = $1;"

	var err error
	_, err = tx.Exec(updateCustomer, customer.Id_Customer, customer.Name, customer.No_Telp, customer.Alamat)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successfully Update Data")
	}

	return nil

}

func DeleteCustomer(id string, tx *sql.Tx) error {
	exists, err := isCustomerExists(id, tx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Pelanggan dengan ID %s tidak ditemukan", id)
	}

	hasTransactions, err := hasCustomerTransactions(id, tx)
	if err != nil {
		return err
	}
	if hasTransactions {
		return fmt.Errorf("Pelanggan dengan ID %s memiliki transaksi terkait dan tidak dapat dihapus", id)
	}

	query := "DELETE FROM mst_customers WHERE id_customer = $1;"
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func isCustomerExists(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM mst_customers WHERE id_customer = $1;"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func hasCustomerTransactions(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM trx_order WHERE customer_id = $1;"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
