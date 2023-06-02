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
	"time"
)

func TransaksiMasterOrder(order model.Order) {
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
		fmt.Println("1. Create Order")
		fmt.Println("2. View All Orders")
		fmt.Println("3. View Order by ID")
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
			fmt.Print("Masukkan ID Order: ")
			idOrder, _ := reader.ReadString('\n')

			fmt.Print("Masukkan ID Customer: ")
			customerId, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Tanggal Masuk (YYYY-MM-DD): ")
			tanggalMasukStr, _ := reader.ReadString('\n')
			tanggalMasukStr = strings.TrimSpace(tanggalMasukStr)
			tanggalMasukArr := strings.Split(tanggalMasukStr, "-")
			tahunMasuk, _ := strconv.ParseInt(tanggalMasukArr[0], 10, 64)
			bulanMasuk, _ := strconv.ParseInt(tanggalMasukArr[1], 10, 64)
			hariMasuk, _ := strconv.ParseInt(tanggalMasukArr[2], 10, 64)
			tanggalMasuk := time.Date(int(tahunMasuk), time.Month(bulanMasuk), int(hariMasuk), 0, 0, 0, 0, time.UTC)

			fmt.Print("Masukkan Tanggal Keluar (YYYY-MM-DD): ")
			tanggalKeluarStr, _ := reader.ReadString('\n')
			tanggalKeluarStr = strings.TrimSpace(tanggalKeluarStr)
			tanggalKeluarArr := strings.Split(tanggalKeluarStr, "-")
			tahunKeluar, _ := strconv.ParseInt(tanggalKeluarArr[0], 10, 64)
			bulanKeluar, _ := strconv.ParseInt(tanggalKeluarArr[1], 10, 64)
			hariKeluar, _ := strconv.ParseInt(tanggalKeluarArr[2], 10, 64)
			tanggalKeluar := time.Date(int(tahunKeluar), time.Month(bulanKeluar), int(hariKeluar), 0, 0, 0, 0, time.UTC)

			newOrder := model.Order{
				Id_Order:       strings.TrimSpace(idOrder),
				Customer_Id:    strings.TrimSpace(customerId),
				Tanggal_Masuk:  tanggalMasuk,
				Tanggal_Keluar: tanggalKeluar,
			}

			err = controller.CreateOrder(newOrder, tx)
			if err != nil {
				fmt.Println("Gagal membuat pesanan:", err)
				continue
			}

			fmt.Println("Pesanan berhasil dibuat.")

			err = tx.Commit()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		case 2:
			orders, err := controller.GetAllOrders(tx)
			if err != nil {
				fmt.Println("Error:", err)
			}
			for _, order := range orders {
				fmt.Println(order)
			}

		case 3:
			fmt.Print("Masukkan ID Order yang akan ditampilkan: ")
			idOrderByID, _ := reader.ReadString('\n')

			orderByID, err := controller.GetOrderById(strings.TrimSpace(idOrderByID), tx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Order by ID:", orderByID)
			}

		case 0:
			fmt.Println("Keluar dari menu Order")
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
