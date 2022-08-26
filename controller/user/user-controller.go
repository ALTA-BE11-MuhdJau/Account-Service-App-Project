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
	// var auth []entities.User
	// err := db.QueryRow("select id from user where telp = ? AND password = ?", inpTelp, inpPass).Scan(&auth)
	// if err != nil {
	// 	return nil, err, "Error select"
	// }
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
