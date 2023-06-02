package controller

import (
	"database/sql"
	"enigma_laundry/model"
	"errors"
	"fmt"
)

func CreateOrderDetail(orderDetail model.OrderDetail, tx *sql.Tx) error {
	if orderDetail.Id_Order_Detail == "" || orderDetail.Order_Id == "" || orderDetail.Layanan_Id == "" {
		return errors.New("Semua field harus diisi")
	}

	exists, err := isOrderDetailExists(orderDetail.Id_Order_Detail, tx)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("Order Detail dengan ID %s sudah ada", orderDetail.Id_Order_Detail)
	}

	exists, err = isLayananExists(orderDetail.Layanan_Id, tx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Layanan dengan ID %s tidak ditemukan", orderDetail.Layanan_Id)
	}

	if orderDetail.Quantity <= 0 {
		return errors.New("Quantity harus lebih dari 0")
	}

	query := "INSERT INTO trx_order_detail (id_order_detail, order_id, layanan_id, quantity) VALUES ($1, $2, $3, $4);"
	_, err = tx.Exec(query, orderDetail.Id_Order_Detail, orderDetail.Order_Id, orderDetail.Layanan_Id, orderDetail.Quantity)
	if err != nil {
		return err
	} else {
		fmt.Println("Successfully Insert Data")
	}

	return nil
}

func isOrderDetailExists(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM trx_order_detail WHERE id_order_detail = $1;"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func isLayananOderExists(id string, tx *sql.Tx) (bool, error) {
	query := "SELECT COUNT(*) FROM mst_layanan WHERE id_layanan = $1;"
	var count int
	err := tx.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetAllOrderDetail(tx *sql.Tx) ([]model.OrderDetail, error) {
	query := "SELECT id_order_detail, order_id, layanan_id, quantity FROM trx_order_detail;"
	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderDetails := []model.OrderDetail{}
	for rows.Next() {
		var orderDetail model.OrderDetail
		err := rows.Scan(&orderDetail.Id_Order_Detail, &orderDetail.Order_Id, &orderDetail.Layanan_Id, &orderDetail.Quantity)
		if err != nil {
			return nil, err
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	return orderDetails, nil
}

func GetOrderDetailById(id string, tx *sql.Tx) (*model.OrderDetail, error) {

	queryOrderDetailById := "SELECT id_order_detail, order_id, layanan_id, quantity FROM trx_order_detail WHERE id_order_detail = $1;"
	row := tx.QueryRow(queryOrderDetailById, id)

	orderDetail := &model.OrderDetail{}
	err := row.Scan(&orderDetail.Id_Order_Detail, &orderDetail.Order_Id, &orderDetail.Layanan_Id, &orderDetail.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Order Detail with ID %s not found", id)
		}
		return nil, err
	}

	return orderDetail, nil

}
