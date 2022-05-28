package db

import (
	"log"
	"recommendersystem/internal/storage"
)

func (d *DB) GetTickets(ticket storage.TicketsRequest) ([]storage.Ticket, error) {
	var id string
	stmt := `select * from tickets where "from"=$1 and "to"=$2 and "date"=$3`
	rows, err := d.Conn.Query(stmt, ticket.From, ticket.To, ticket.Date)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var res []storage.Ticket
	var t storage.Ticket
	for rows.Next() {
		err = rows.Scan(&id, &t.From, &t.To, &t.Date, &t.DepartureTime, &t.ArrivalTime, &t.Transfer, &t.TransferDepartureTime, &t.TransferTime)
		if err != nil {
			log.Println(err)
		}
		res = append(res, t)
	}
	rows.Close()

	return res, nil
}
