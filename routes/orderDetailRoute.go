package routes

import (
	"bufio"
	"enigma_laundry/config"
	"enigma_laundry/controller"
	"enigma_laundry/model"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TransaksiMasterOrderDetail(orderDetail model.OrderDetail) {
	// Mengkoneksi ke Database
	db := config.ConnectDb()
	defer db.Close()

	// Membuat transaksi
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	reader := bufio.NewReader(os.Stdin)

	for {
		// Menampilkan pilihan menu
		fmt.Println("Menu:")
		fmt.Println("1. Create Order Detail")
		fmt.Println("2. View All Orders Detail")
		fmt.Println("3. View Order Detail by ID")
		fmt.Println("0. Exit")
		fmt.Print("Pilih menu: ")

		// Membaca input pilihan menu dari pengguna
		menuStr, _ := reader.ReadString('\n')
		menu, err := strconv.Atoi(strings.TrimSpace(menuStr))
		if err != nil {
			fmt.Println("Input tidak valid")
			continue
		}

		switch menu {
		case 1:
			fmt.Print("Masukkan ID Order Detail: ")
			idOrderDetail, _ := reader.ReadString('\n')

			fmt.Print("Masukkan ID Order: ")
			orderId, _ := reader.ReadString('\n')

			fmt.Print("Masukkan ID Layanan: ")
			layananId, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Jumlah Order: ")
			quantityStr, _ := reader.ReadString('\n')
			quantity, _ := strconv.Atoi(strings.TrimSpace(quantityStr))

			newOrderDetail := model.OrderDetail{
				Id_Order_Detail: strings.TrimSpace(idOrderDetail),
				Order_Id:        strings.TrimSpace(orderId),
				Layanan_Id:      strings.TrimSpace(layananId),
				Quantity:        quantity,
			}

			err = controller.CreateOrderDetail(newOrderDetail, tx)
			if err != nil {
				fmt.Println("Gagal Menambahkan Detail Laundry:", err)
				continue
			}

			fmt.Println("Oder Detail berhasil dibuat.")

			err = tx.Commit()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		case 2:
			orderDetails, err := controller.GetAllOrderDetail(tx)
			if err != nil {
				fmt.Println("Error:", err)
			}
			for _, orderDetail := range orderDetails {
				fmt.Println(orderDetail)
			}

		case 3:
			fmt.Print("Masukkan ID Order Detail yang akan ditampilkan: ")
			idOrderDetailByID, _ := reader.ReadString('\n')

			orderDetailByID, err := controller.GetOrderDetailById(strings.TrimSpace(idOrderDetailByID), tx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Order by ID:", orderDetailByID)
			}

		case 0:
			fmt.Println("Keluar dari menu Order detail")
			break

		default:
			fmt.Println("Menu tidak valid")
		}

		if menu == 0 {
			break
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
