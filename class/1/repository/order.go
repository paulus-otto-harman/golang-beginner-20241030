package repository

import (
	_interface "20241030/class/1/repository/interface"
	"database/sql"
)

type Order struct {
	Db *sql.DB
}

func Get(db *sql.DB) _interface.OrderRepository {
	return &Order{Db: db}
}

func getRows(order *Order, query string) *sql.Rows {
	rows, err := order.Db.Query(query)
	if err != nil {
		return nil
	}
	return rows
}

func (order *Order) Get1(query string) (interface{}, error) {
	rows := getRows(order, query)
	defer rows.Close()

	var hasil []struct {
		bulan  string
		jumlah int
	}
	for rows.Next() {
		baris := struct {
			bulan  string
			jumlah int
		}{}
		err := rows.Scan(&baris.bulan, &baris.jumlah)
		hasil = append(hasil, baris)
		if err != nil {
			return nil, err
		}
	}
	return hasil, nil
}

func (order *Order) Get2(query string) (interface{}, error) {
	rows := getRows(order, query)
	defer rows.Close()

	var hasil []struct {
		bulan  string
		nama   string
		jumlah int
	}
	for rows.Next() {
		baris := struct {
			bulan  string
			nama   string
			jumlah int
		}{}
		err := rows.Scan(&baris.bulan, &baris.nama, &baris.jumlah)
		hasil = append(hasil, baris)

		if err != nil {
			return nil, err
		}
	}
	return hasil, nil
}

func (order *Order) Get3(query string) (interface{}, error) {
	rows := getRows(order, query)
	defer rows.Close()

	var hasil []struct {
		area   string
		jumlah int
	}
	for rows.Next() {
		baris := struct {
			area   string
			jumlah int
		}{}
		err := rows.Scan(&baris.area, &baris.jumlah)
		hasil = append(hasil, baris)

		if err != nil {
			return nil, err
		}
	}
	return hasil, nil
}

func (order *Order) Get4(query string) (interface{}, error) {
	rows := getRows(order, query)
	defer rows.Close()

	var hasil []struct {
		jam    string
		jumlah int
	}
	for rows.Next() {
		baris := struct {
			jam    string
			jumlah int
		}{}
		err := rows.Scan(&baris.jam, &baris.jumlah)
		hasil = append(hasil, baris)

		if err != nil {
			return nil, err
		}
	}
	return hasil, nil
}

func (order *Order) Get5(query string) (int, error) {
	row := order.Db.QueryRow(query)
	var jumlah int
	err := row.Scan(&jumlah)
	if err != nil {
		return 0, err
	}
	return jumlah, nil
}
