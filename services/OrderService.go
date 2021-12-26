package services

import (
	"RestApiWithRedisAndRabbitMQ/config"
	"RestApiWithRedisAndRabbitMQ/models"
	"database/sql"
	"fmt"
	"log"
)


func InsertOrder(order models.Order) int64{
	connection:= config.InitDB() // create the postgres db connection
	defer connection.Close() // close connection
	sqlStatement := `INSERT into orders (code,description, price) VALUES ($1, $2, $3) `
	var status int64  = 1
	_, err := connection.Exec(sqlStatement, order.Code, order.Description, order.Price)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return 0
	}
	return status
}

func GetOrder(id int64) (models.Order, error){
	connection:= config.InitDB() // create the postgres db connection
	defer connection.Close() // close connection
	var order models.Order
	sqlStatement := `SELECT * FROM orders WHERE id=$1`
	row := connection.QueryRow(sqlStatement, id)
	err := row.Scan(&order.Id, &order.Code, &order.Description, &order.Price, &order.Created_at, &order.Updated_at)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return order, nil
	case nil:
		return order, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return order, err
}

// get all orders from the DB by its userid
func GetAllOrders() ([]models.Order, error) {
	connection:= config.InitDB() // create the postgres db connection
	defer connection.Close() // close connection
	var orders []models.Order
	// create the select sql query
	sqlStatement := `SELECT * FROM orders`
	// execute the sql statement
	rows, err := connection.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	// close the statement
	defer rows.Close()
	// iterate over the rows
	for rows.Next() {
		var order models.Order
		// unmarshal the row object to user
		err = rows.Scan(&order.Id, &order.Code, &order.Description, &order.Price, &order.Created_at, &order.Updated_at)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// append the user in the users slice
		orders = append(orders, order)
	}
	// return empty user on error
	return orders, err
}

// update order in the DB
func UpdateOrder(id int64, order models.Order) int64 {
	connection:= config.InitDB() // create the postgres db connection
	defer connection.Close() // close connection
	// create the update sql query
	sqlStatement := `UPDATE orders SET code=$2, description=$3, price=$4 WHERE id=$1`
	// execute the sql statement
	res, err := connection.Exec(sqlStatement, id, order.Code, order.Description, order.Price)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}

// delete order in the DB
func DeleteOrder(id int64) int64 {
	connection:= config.InitDB() // create the postgres db connection
	defer connection.Close() // close connection
	// create the delete sql query
	sqlStatement := `DELETE FROM orders WHERE id=$1`
	// execute the sql statement
	res, err := connection.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
