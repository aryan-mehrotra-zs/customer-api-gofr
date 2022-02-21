package customer

const (
	create = "INSERT INTO customers(name) VALUES ($1) RETURNING id;"
	get    = "SELECT * FROM customers WHERE id=$1;"
	update = "UPDATE customers SET name=($1) WHERE id=($2)"
	delete = "DELETE FROM customers WHERE id=($1)"
)
