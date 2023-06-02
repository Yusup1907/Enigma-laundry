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

func MenuMasterLayanan(layanan model.Layanan) {
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
		fmt.Println("1. Add Layanan")
		fmt.Println("2. Update Layanan")
		fmt.Println("3. Delete Layanan")
		fmt.Println("4. View All Layanan")
		fmt.Println("5. View Layanan by ID")
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
			fmt.Print("Masukkan ID Layanan: ")
			idLayanan, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Nama Layanan: ")
			namaLayanan, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Harga Layanan: ")
			hargaLayananStr, _ := reader.ReadString('\n')
			hargaLayanan, _ := strconv.Atoi(strings.TrimSpace(hargaLayananStr))

			fmt.Print("Masukkan Satuan Layanan: ")
			satuanLayanan, _ := reader.ReadString('\n')

			layanan := model.Layanan{
				Id_Layanan:   strings.TrimSpace(idLayanan),
				Nama_Layanan: strings.TrimSpace(namaLayanan),
				Harga:        hargaLayanan,
				Satuan:       strings.TrimSpace(satuanLayanan),
			}
			err = controller.CreateLayanan(layanan, tx)
			if err != nil {
				fmt.Println("Gagal menyimpan layanan:", err)
				continue
			}

			fmt.Println("Layanan berhasil disimpan.")
			err = tx.Commit()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

		case 2:
			fmt.Print("Masukkan ID Layanan yang akan diupdate: ")
			idLayananUpdate, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Nama Layanan baru: ")
			namaLayananUpdate, _ := reader.ReadString('\n')

			fmt.Print("Masukkan Harga Layanan baru: ")
			hargaLayananStr, _ := reader.ReadString('\n')
			hargaLayanan, _ := strconv.Atoi(strings.TrimSpace(hargaLayananStr))

			fmt.Print("Masukkan Satuan Layanan baru: ")
			satuanLayananUpdate, _ := reader.ReadString('\n')

			layananUpdate := model.Layanan{
				Id_Layanan:   strings.TrimSpace(idLayananUpdate),
				Nama_Layanan: strings.TrimSpace(namaLayananUpdate),
				Harga:        hargaLayanan,
				Satuan:       strings.TrimSpace(satuanLayananUpdate),
			}

			err = controller.UpdateLayanan(layananUpdate, tx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Layanan berhasil diupdate")
			}

			err = tx.Commit()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

		case 3:
			fmt.Print("Masukkan ID Layanan yang akan dihapus: ")
			idLayananDelete, _ := reader.ReadString('\n')

			err = controller.DeleteLayanan(strings.TrimSpace(idLayananDelete), tx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Layanan berhasil dihapus")
			}

			err = tx.Commit()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

		case 4:
			layanans := controller.GetAllLayanan()
			for _, layanan := range layanans {
				fmt.Println(layanan)
			}

		case 5:
			fmt.Print("Masukkan ID Layanan yang akan ditampilkan: ")
			idLayananByID, _ := reader.ReadString('\n')

			layananByID, err := controller.GetLayananById(strings.TrimSpace(idLayananByID), tx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Layanan by ID:", layananByID)
			}

		case 0:
			fmt.Println("Keluar dari menu Master Layanan")
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
