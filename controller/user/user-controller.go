package user

import (
	"be11/project-app/entities"
	"database/sql"
)

func AddUser(db *sql.DB, addUser entities.User) (int, error) {
	stmt, errPrep := db.Prepare("insert into user (name, dob, gender, telp, password, sisa_saldo) values (?, ?, ?, ?, ?, ?)")
	if errPrep != nil {
		return -1, errPrep
	}
	result, errExec := stmt.Exec(addUser.Name, addUser.DoB, addUser.Gender, addUser.Telp, addUser.Password, addUser.SisaSaldo)
	if errExec != nil {
		return -1, nil
	} else {
		row, errRow := result.RowsAffected()
		if errRow != nil {
			return 0, nil
		}
		return int(row), nil
	}
}

func LoginUser(db *sql.DB, inpTelp, inpPass string) ([]entities.User, error) {
	// oldData := entities.User{}
	// err := db.QueryRow("SELECT no_rekening, nomor_telepon, password FROM users WHERE nomor_telepon = ?", newUser.No_telepon).Scan(&oldUser.No_rekening, &oldUser.No_telepon, &oldUser.Password)
	// if err != nil {
	// 	return -1, "", "", "", err
	// }
	// if newUser.No_telepon == oldUser.No_telepon && newUser.Password == oldUser.Password {
	// 	login = "login sukses"
	// 	fmt.Println(login)
	// } else {
	// 	login = "No Telepon Atau Password Salah"
	// 	fmt.Println(login)
	// }
	// return 0, login, oldUser.No_rekening, oldUser.No_telepon, nil

	res, err := db.Query("select * from user where telp = ?", inpTelp)
	if err != nil {
		return nil, err
	}
	dataUser := []entities.User{}
	for res.Next() {
		var rowUser entities.User
		err := res.Scan(&rowUser.ID, &rowUser.Name, &rowUser.DoB, &rowUser.Gender, &rowUser.Telp, &rowUser.Password, &rowUser.SisaSaldo, &rowUser.CreatedAt, &rowUser.UpdatedAt)
		if err != nil {
			return nil, err
		}
		dataUser = append(dataUser, rowUser)
		if rowUser.Password != inpPass {
			return nil, sql.ErrNoRows
		} else {
			return dataUser, nil
		}
	}
	return nil, nil
}

// func ReadUser(db *sql.DB, readUser entities.User) (entities.User, error) {
// 	prevData := entities.User{}
// 	err := db.QueryRow("select id, name, dob, gender, telp, password, sisa_saldo from user where id = ?", readUser.ID).Scan(&prevData.ID, &prevData.Name, &prevData.DoB, &prevData.Gender, &prevData.Telp, &prevData.Password, &prevData.SisaSaldo)
// 	if err != nil {
// 		return prevData, err
// 	}
// 	return prevData, nil
// }

func UpdateUser(db *sql.DB, updateUser entities.User) (int, error) {
	prevDataUser := entities.User{}

	err := db.QueryRow("select * from user where telp = ?", updateUser.Telp).Scan(&updateUser.ID, &prevDataUser.Name, &prevDataUser.DoB, &prevDataUser.Gender, &prevDataUser.Telp, &prevDataUser.Password, &prevDataUser.SisaSaldo, &prevDataUser.CreatedAt, &prevDataUser.UpdatedAt)
	if err != nil {
		return -1, err
	}
	if updateUser.Name == "" {
		updateUser.Name = prevDataUser.Name
	}
	if updateUser.DoB == "" {
		updateUser.DoB = prevDataUser.DoB
	}
	if updateUser.Gender == "" {
		updateUser.Gender = prevDataUser.Gender
	}
	if updateUser.Telp == "" {
		updateUser.Telp = prevDataUser.Telp
	}
	if updateUser.Password == "" {
		updateUser.Password = prevDataUser.Password
	}

	stmt, errPrep := db.Prepare("update user set name = ?, dob = ?, gender = ?, telp = ?, password = ? where id = ?")
	if errPrep != nil {
		return -1, errPrep
	}
	res, errStmt := stmt.Exec(updateUser.Name, updateUser.DoB, updateUser.Gender, updateUser.Telp, updateUser.Password)
	if errStmt != nil {
		return -1, errStmt
	} else {
		rowUser, err := res.RowsAffected()
		if err != nil {
			return 0, err
		}
		return int(rowUser), err
	}
}

func DeleteUser(db *sql.DB, inpTelp string) (int, error) {
	stmt, err := db.Prepare("delete from user where telp = ?")
	if err != nil {
		return -1, err
	}
	result, err := stmt.Exec(&inpTelp)
	if err != nil {
		return -1, err
	} else {
		rowUser, err := result.RowsAffected()
		if err != nil {
			return 0, err
		}
		return int(rowUser), err
	}
}

func GetOtherUser(db *sql.DB, inpTelp string) ([]entities.User, error) {
	result, err := db.Query("select id, name, dob, gender from user where telp = ?", inpTelp)
	if err != nil {
		return nil, err
	}
	var otherUser []entities.User
	for result.Next() {
		var rowUser entities.User
		err := result.Scan(&rowUser.ID, &rowUser.Name, &rowUser.DoB, &rowUser.Gender)
		if err != nil {
			return nil, err
		}
		otherUser = append(otherUser, rowUser)
	}
	return otherUser, nil
}
