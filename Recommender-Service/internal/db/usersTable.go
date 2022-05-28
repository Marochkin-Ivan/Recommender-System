package db

import (
	"log"
	"recommendersystem/internal/storage"
)

func (d *DB) RegistrationUser(u storage.UserInfo) (string, bool, error) {
	stmt := `select uid from users where log=$1`
	rows, err := d.Conn.Query(stmt, u.Login)
	if err != nil {
		return "", false, err
	}
	if rows.Err() != nil {
		return "", false, rows.Err()
	}

	if rows.Next() {
		return "", true, nil
	}
	rows.Close()

	uid := GenerateUID()
	stmt = `insert into users(uid,log,pas) values($1,$2,$3)`
	_, err = d.Conn.Exec(stmt, uid, u.Login, u.Password)
	if err != nil {
		return "", false, err
	}

	return uid, false, nil
}

func (d *DB) AuthorizationUser(u storage.UserInfo) (string, bool, error) {
	stmt := `select uid from users where log=$1 and pas=$2`
	rows, err := d.Conn.Query(stmt, u.Login, u.Password)
	if err != nil {
		return "", false, err
	}
	if rows.Err() != nil {
		return "", false, rows.Err()
	}

	defer rows.Close()
	var uid string
	if rows.Next() {
		err = rows.Scan(&uid)
		if err != nil {
			log.Println(err)
			return "", false, err
		}
	} else {
		return "", false, nil
	}

	return uid, true, nil
}

func (d *DB) GetUserName(id string) (string, error) {
	var name string
	stmt := `select log from users where uid=$1`
	rows, err := d.Conn.Query(stmt, id)
	if err != nil {
		return "", err
	}
	if rows.Err() != nil {
		return "", rows.Err()
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Println(err)
			return "", err
		}
	} else {
		return "", nil
	}

	return name, nil
}
