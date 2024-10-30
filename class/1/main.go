package main

import (
	"20241030/class/1/service"
	"20241030/database"
)

func main() {
	db := database.DbOpen("20241029c")
	defer db.Close()

	// contoh menambahkan customer
	//err := service.NewCustomer(db, "customer enam", "customer6", "password")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// contoh menambahkan driver
	//err := service.NewDriver(db, "driver empat", "driver4", "password")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// contoh menambahkan order

	//err := service.NewOrder(db, )
	//if err != nil {
	//	log.Fatal(err)
	//}

	//contoh query order
	service.Get1(db, "SELECT to_char(created_at,'Mon') AS bulan, count(*) AS total_per_bulan FROM orders GROUP BY bulan")

	service.Get2(db, "SELECT to_char(created_at,'Mon') bulan, customers.name, count(*) jumlah FROM orders JOIN customers ON orders.customer_id=customers.id GROUP BY bulan, customers.name ORDER BY bulan, jumlah DESC")

	service.Get3(db, "SELECT areas.name, count(*) FROM orders JOIN areas ON orders.area_id=areas.id GROUP BY areas.name ORDER BY 2 DESC")

	service.Get4(db, "SELECT date_part('hour', created_at) jam, count(*) FROM orders GROUP BY jam ORDER BY 2 DESC")

	service.Get5(db, "SELECT count(*) FROM sessions WHERE expired_at IS NULL")

	service.Get2(db, "SELECT  to_char(orders.created_at,'Mon') bulan, drivers.name, count(*) FROM orders JOIN drivers ON orders.driver_id=drivers.id GROUP BY bulan,drivers.name ORDER BY bulan,3 DESC")

}
