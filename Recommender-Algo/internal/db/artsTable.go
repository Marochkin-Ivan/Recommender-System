package db

import (
	"log"
	"recommenderalgo/internal/algoStorage"
	"strconv"
)

func (d *DB) GetArtsForRecommend() (map[string]map[string]float64, error) {
	var id string
	stmt := `select * from points`
	rows, err := d.Conn.Query(stmt)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	res := make(map[string]map[string]float64)
	var p algoStorage.Points

	for rows.Next() {
		err = rows.Scan(&id, &p.Name, &p.Art, &p.Points)
		if err != nil {
			log.Println(err)
		}
		if res[p.Name] == nil {
			res[p.Name] = make(map[string]float64)
		}
		res[p.Name][p.Art] = p.Points
	}
	rows.Close()

	return res, nil
} // return big map...

func (d *DB) GetArtsByCity(transfer algoStorage.RecommendRequest) (map[string]string, error) {
	stmt := `select art, intime from arts where city=$1`
	rows, err := d.Conn.Query(stmt, transfer.City)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	res := make(map[string]string)
	var a algoStorage.Arts
	it, _ := strconv.Atoi(transfer.InTime)

	for rows.Next() {
		err = rows.Scan(&a.Art, &a.InTime)
		if err != nil {
			log.Println(err)
		}

		tmp, _ := strconv.Atoi(a.InTime)
		if tmp <= it {
			res[a.Art] = a.InTime
		}
	}
	rows.Close()

	return res, nil
} // return internal id
