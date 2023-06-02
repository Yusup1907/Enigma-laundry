package controller

import (
	"database/sql"
	"enigma_laundry/model"
	"errors"
	"fmt"
)

func CreateOrder(order model.Order, tx *sql.Tx) error {
	if order.Id_Order == "" || order.Customer_Id == "" || order.Tanggal_Masuk.IsZero() || order.Tanggal_Keluar.IsZero() {
		return errors.New("Semua field harus diisi")
	}

	exists, err := isOrderExists(order.Id_Order, tx)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("Order dengan ID %s sudah ada", order.Id_Order)
	}

	if order.Tanggal_Keluar.Before(order.Tanggal_Masuk) {
		return errors.New("Tanggal Keluar harus lebih besar atau sama dengan Tanggal Masuk")
	}

	exists, err = isCustomerOderExists(order.Customer_Id, tx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Pelanggan dengan ID %s tidak ditemukan", order.Customer_Id)
	}

	query := "INSERT INTO trx_order (id_order, customer_id, tanggal_masuk, tanggal_keluar) VALUES ($1, $2, $3, $4);"
	_, err = tx.Exec(query, order.Id_Order, order.Customer_Id, order.Tanggal_Masuk, order.Tanggal_Keluar)
	if err != nil {
		return err
	} else {
		fmt.Println("successfully Insert Data")
	}

	return nil
}

func isOrderExists(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM trx_order WHERE id_order = $1;"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func isCustomerOderExists(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM mst_customers WHERE id_customer = $1;"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetAllOrders(tx *sql.Tx) ([]model.Order, error) {
	query := "SELECT id_order, customer_id, tanggal_masuk, tanggal_keluar FROM trx_order;"
	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []model.Order{}
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.Id_Order, &order.Customer_Id, &order.Tanggal_Masuk, &order.Tanggal_Keluar)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderById(id string, tx *sql.Tx) (*model.Order, error) {

	queryOrderById := "SELECT id_order, customer_id, tanggal_masuk, tanggal_keluar FROM trx_order WHERE id_order = $1;"
	row := tx.QueryRow(queryOrderById, id)

	order := &model.Order{}
	err := row.Scan(&order.Id_Order, &order.Customer_Id, &order.Tanggal_Masuk, &order.Tanggal_Keluar)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Order with ID %s not found", id)
		}
		return nil, err
	}

	return order, nil

}
