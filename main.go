package main

import (
	"be11/project-app/config"
	"be11/project-app/controller/user"
	"be11/project-app/entities"
	"fmt"
)

func main() {
	db := config.DBconn()

	defer db.Close()

	var menu int
	for menu != 99 {
		fmt.Println("-------------------------------")
		fmt.Print("📌 Menu Account Service App:\n\n1. Add Account (Register)\n2. Login Account\n99. Exit\n\n")
		fmt.Print("📌 Masukkan pilihan anda: ")
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			{
				addUser := entities.User{}

				fmt.Println("-------------------------------")
				fmt.Println("Nama User Anda:")
				fmt.Scanln(&addUser.Name)
				fmt.Println("Tanggal Lahir:")
				fmt.Scanln(&addUser.DoB)
				fmt.Println("Jenis Kelamin:")
				fmt.Scanln(&addUser.Gender)
				fmt.Println("Nomor Telepon:")
				fmt.Scanln(&addUser.Telp)
				fmt.Println("Input Password:")
				fmt.Scanln(&addUser.Password)
				fmt.Println("Saldo Awal:")
				fmt.Scanln(&addUser.SisaSaldo)

				rowAffected, err := user.AddUser(db, addUser)
				if err != nil {
					fmt.Println("Error insert data.", err)
				} else {
					if rowAffected == 0 {
						fmt.Println("Gagal insert data.")

					} else {
						fmt.Println("Insert sukses, Row affected:", rowAffected)
					}
				}
			}
		case 2:
			{
				loginUser := entities.User{}

				fmt.Println("-------------------------------")
				fmt.Print("Nomor Telepon: ")
				fmt.Scanln(&loginUser.Telp)
				fmt.Print("Password: ")
				fmt.Scanln(&loginUser.Password)

				_, err := user.LoginUser(db, loginUser.Telp, loginUser.Password)
				if err != nil {
					fmt.Println("Gagal Login!", err)
				} else {
					// Menu setelah login:
					var secondMenu int
					for secondMenu != 99 {
						fmt.Println(secondMenu)
						fmt.Println("-------------------------------")
						fmt.Print("📌 Menu Account Service App:\n\n1. Profile Anda\n2. Update Account\n3. Delete Account\n4. Top-Up\n5. Transfer\n6. History Top-Up\n7. Histroy Transfer\n8. Check Others\n99. Exit\n\n")
						fmt.Print("📌 Masukkan pilihan anda: ")
						fmt.Scanln(&secondMenu)
						switch secondMenu {
						case 1:
							{

							}
						case 2:
							{
								updateUser := entities.User{}
								fmt.Println(updateUser)

								fmt.Println("-------------------------------")
								fmt.Println("Nama User Anda:")
								fmt.Scanln(&updateUser.Name)
								fmt.Println("Tanggal Lahir:")
								fmt.Scanln(&updateUser.DoB)
								fmt.Println("Jenis Kelamin:")
								fmt.Scanln(&updateUser.Gender)
								fmt.Println("Nomor Telepon:")
								fmt.Scanln(&updateUser.Telp)
								fmt.Println("Input Password:")
								fmt.Scanln(&updateUser.Password)
								fmt.Println(updateUser)

								rowAffect, err := user.UpdateUser(db, updateUser)
								if err != nil {
									fmt.Println("Error update data.", err)
								} else {
									if rowAffect == 0 {
										fmt.Println("Tidak ada data yang diubah!")
									} else {
										fmt.Println("Update sukses, Row affected =", rowAffect)
									}
								}
							}
						case 3:
							{
								deleteUser := entities.User{}

								fmt.Println("-------------------------------")
								fmt.Println("(Izinkan menghapus, dengan input nomor telepon anda sekarang)")
								fmt.Print("Nomor Telepon: ")
								fmt.Scanln(&deleteUser.Telp)

								res, err := user.DeleteUser(db, deleteUser.Telp)
								if err != nil {
									fmt.Println("Gagal menghapus!", err)
								} else if res >= 0 {
									fmt.Println("Berhasil menghapus!")
									secondMenu = 99
									break
								}
							}
						case 4:
							{

							}
						case 5:
							{

							}
						case 6:
							{

							}
						case 7:
							{

							}
						case 8:
							{
								var otherUser string
								fmt.Println("Masukkan nomor telepon:")
								fmt.Scanln(&otherUser)

								result, err := user.GetOtherUser(db, otherUser)
								if err != nil {
									fmt.Println("Error read data from database.", err)
								} else {
									for _, col := range result {
										fmt.Println("   ID:", col.ID)
										fmt.Println("   Name:", col.Name)
										fmt.Println("   Birth:", col.DoB)
										fmt.Println("   Gender:", col.Gender)
										fmt.Println("-------------------------------")
									}
								}
							}
						case 99:
							{
								break
							}
						default:
							{
								fmt.Println("-------------------------------")
								fmt.Println("Menu yang anda pilih tidak ada!")
								continue
							}
						}
					}

				}
			}
		case 99:
			{
				break
			}
		default:
			{
				fmt.Println("-------------------------------")
				fmt.Println("Menu yang anda pilih tidak ada!")
				continue
			}
		}
	}
	fmt.Println("--------------------------------")
	fmt.Println("Terima kasih telah bertransaksi!")
	fmt.Println("--------------------------------")
}
