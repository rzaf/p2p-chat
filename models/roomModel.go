package models

import (
	"errors"
	"fmt"
	"math/rand"
)

type Room struct {
	Id       int64
	Name     string
	Addr     string
	Port     string
	Secrete  string
	Uuid     string
	UserUuid string
}

func MigrateRooms() error {
	if Db == nil {
		return nil
	}
	query := `
	create table if not exists rooms(
		id integer primary key autoincrement,
		Name varchar(128) not null unique,
		Addr varchar(128) not null,
		Port varchar(128) not null,
		Secrete varchar(128) not null,
		Uuid varchar(128) not null,
		UserUuid varchar(128) not null
	);
	`
	_, err := Db.Exec(query)
	return err
}

func InsertRoom(r Room) (*Room, error) {
	if Db == nil {
		r.Id = rand.Int63()
		return &r, nil
	}
	s := "insert into rooms (Name,Addr,Port,Secrete,Uuid,UserUuid) values (?,?,?,?,?,?);"

	res, err := Db.Exec(s, r.Name, r.Addr, r.Port, r.Secrete, r.Uuid, r.UserUuid)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	r.Id = id
	return &r, nil
}

func GetAllRooms() ([]*Room, error) {
	if Db == nil {
		return nil, nil
	}
	// s := "select (Name,Addr,Port,Secrete,Uuid) from rooms"
	s := "select * from rooms;"
	rows, err := Db.Query(s)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var AllRooms []*Room
	for rows.Next() {
		var r Room
		err2 := rows.Scan(&r.Id, &r.Name, &r.Addr, &r.Port, &r.Secrete, &r.Uuid, &r.UserUuid)
		if err2 != nil {
			return nil, err2
		}
		AllRooms = append(AllRooms, &r)
	}
	return AllRooms, nil
}

func GetRoomById(id int64) (*Room, error) {
	if Db == nil {
		return nil, nil
	}
	s := "select (Name,Addr,Port,Secrete,Uuid,UserUuid) from rooms WHERE id=?;"

	row := Db.QueryRow(s, id)
	if row == nil {
		return nil, errors.New("not found")
	}
	var r Room
	err := row.Scan(&r.Id, &r.Name, &r.Addr, &r.Port, &r.Secrete, &r.Uuid, &r.UserUuid)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func UpdateRoom(id int64, updated Room) error {
	if Db == nil {
		return nil
	}
	s := "update rooms set Name=?,Addr=?,Port=?,Secrete=?,Uuid=?,UserUuid=? WHERE id=?;"

	r, err := Db.Exec(s, updated.Name, updated.Addr, updated.Port, updated.Secrete, updated.Uuid, updated.UserUuid, id)
	if err != nil {
		return err
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func DeleteRoom(id int64) error {
	if Db == nil {
		return nil
	}
	s := "delete from rooms WHERE id=?;"

	r, err := Db.Exec(s, id)
	if err != nil {
		return err
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}
