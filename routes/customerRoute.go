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

func MenuMasterCustomer(customer model.Customer) {
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
		fmt.Println("1. Add Customer")
		fmt.Println("2. Update Customer")
		fmt.Println("3. Delete Customer")
		fmt.Println("4. View All Customers")
		fmt.Println("5. View Customer by ID")
		fmt.Println("0. Keluar")
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
			fmt.Print("Masukkan ID Customer: ")
			idCustomer, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Nama Customer: ")
			namaCustomer, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Nomor Telepon Customer: ")
			teleponCustomer, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Alamat Customer: ")
			alamatCustomer, _ := reader.ReadString('\n')

			customer := model.Customer{
				Id_Customer: strings.TrimSpace(idCustomer),
				Name:        strings.TrimSpace(namaCustomer),
				No_Telp:     strings.TrimSpace(teleponCustomer),
				Alamat:      strings.TrimSpace(alamatCustomer),
			}

			err = controller.CreateCustomer(customer, tx)
			if err != nil {
				fmt.Println("Gagal menyimpan customer:", err)
				continue
			}

			fmt.Println("Customer berhasil disimpan.")
			err = tx.Commit()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

		case 2:
			fmt.Print("Masukkan ID Customer yang akan diupdate: ")
			idCustomerUpdate, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Nama Customer baru: ")
			namaCustomerUpdate, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Nomor Telepon Customer baru: ")
			teleponCustomerUpdate, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Alamat Customer baru: ")
			alamatCustomerUpdate, _ := reader.ReadString('\n')

			customerUpdate := model.Customer{
				Id_Customer: strings.TrimSpace(idCustomerUpdate),
				Name:        strings.TrimSpace(namaCustomerUpdate),
				No_Telp:     strings.TrimSpace(teleponCustomerUpdate),
				Alamat:      strings.TrimSpace(alamatCustomerUpdate),
			}

			err = controller.UpdateCustomer(customerUpdate, tx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Customer berhasil diupdate")
			}

			err = tx.Commit()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

		case 3:
			fmt.Print("Masukkan ID Customer yang akan dihapus: ")
			idCustomerDelete, _ := reader.ReadString('\n')

			err = controller.DeleteCustomer(strings.TrimSpace(idCustomerDelete), tx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Customer berhasil dihapus")
			}

			err = tx.Commit()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

		case 4:
			customers := controller.GetAllCustomer()
			for _, customer := range customers {
				fmt.Println(customer)
			}

		case 5:
			fmt.Print("Masukkan ID Customer yang akan ditampilkan: ")
			idCustomerByID, _ := reader.ReadString('\n')

			customerByID, err := controller.GetCustomerById(strings.TrimSpace(idCustomerByID), tx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Customer by ID:", customerByID)
			}

		case 0:
			fmt.Println("Keluar dari menu Master Customer")
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
