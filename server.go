package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Routes
	e.POST("/employees/add", createUser)
	e.GET("/employees/:id", getUser)
	e.GET("/employees", getAllUser)
	e.PUT("/employees/:id", updateUser)
	e.DELETE("/employees/:id", deleteUser)
	e.GET("/employees/name/:name", getname)

	e.POST("/member/add", createMember)
	e.GET("/member/:id", getMember)
	e.GET("/member", getAllMember)
	e.PUT("/member/:id", updateMember)
	e.DELETE("/member/:id", deleteMember)
	e.POST("/member/login", login)

	e.POST("/products/add", addProduct)
	e.GET("/products", getAllProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.GET("/products/:id", getProduct)
	e.PUT("/products/:id", updateProduct)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func createUser(c echo.Context) error {
	emp := new(Employee)
	if err := c.Bind(emp); err != nil {
		return err
	}

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	sqlStatment := `INSERT INTO employee (name, email) VALUES ($1, $2);`
	_, err = db.Exec(sqlStatment, emp.Name, emp.Email)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "createuser")
}
func getUser(c echo.Context) error {
	id := c.Param("id")

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	var employee Employee

	sqlStatment := `SELECT * FROM employee WHERE id=$1;`
	err = db.QueryRow(sqlStatment, id).Scan(&employee.ID, &employee.Name, &employee.Email)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return c.JSON(http.StatusOK, employee)
}
func getAllUser(c echo.Context) error {
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Succesfully connected! DB")

	sqlStatment := `SELECT * FROM employee;`
	result, err := db.Query(sqlStatment)
	if err != nil {
		fmt.Println(err)
	}
	var employees []Employee

	for result.Next() {
		var data Employee
		err := result.Scan(
			&data.ID,
			&data.Name,
			&data.Email)
		if err != nil {
			return err
		}
		employees = append(employees, data)
	}
	defer result.Close()
	defer db.Close()
	return c.JSON(http.StatusOK, employees)
}
func updateUser(c echo.Context) error {
	id := c.Param("id")
	emp := new(Employee)
	if err := c.Bind(emp); err != nil {
		return err
	}
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	sqlStatment := `UPDATE employee SET name =$1, email=$2 WHERE id =$3;`
	_, err = db.Exec(sqlStatment, emp.Name, emp.Email, id)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "updateuser")
}
func deleteUser(c echo.Context) error {
	id := c.Param("id")

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	sqlStatment := `DELETE FROM employee WHERE id=$1;`
	_, err = db.Exec(sqlStatment, id)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	return c.JSON(http.StatusOK, "deleteuser")
}
func getname(c echo.Context) error {
	name := c.Param("name")

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	var employee Employee
	sqlStatment := `SELECT * FROM employee WHERE name=$1;`
	err = db.QueryRow(sqlStatment, name).Scan(&employee.ID, &employee.Name, &employee.Email)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return c.JSON(http.StatusOK, employee)
}

func createMember(c echo.Context) error {
	mb := new(Member)
	if err := c.Bind(mb); err != nil {
		return err
	}

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	sqlStatment := `INSERT INTO member (username, password, fisrtname, lastname, email, phone) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err = db.Exec(sqlStatment, mb.Username, mb.Password, mb.Fisrtname, mb.Lastname, mb.Email, mb.Phone)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "createmember")
}
func getMember(c echo.Context) error {
	id := c.Param("id")

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	var member Member

	sqlStatment := `SELECT * FROM member WHERE id=$1;`
	err = db.QueryRow(sqlStatment, id).Scan(&member.ID, &member.Username, &member.Password, &member.Fisrtname, &member.Lastname, &member.Email, &member.Phone)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return c.JSON(http.StatusOK, member)
}
func getAllMember(c echo.Context) error {
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Succesfully connected! DB")

	sqlStatment := `SELECT * FROM member;`
	result, err := db.Query(sqlStatment)
	if err != nil {
		fmt.Println(err)
	}
	var member []Member

	for result.Next() {
		var data Member
		err := result.Scan(
			&data.ID,
			&data.Username,
			&data.Password,
			&data.Fisrtname,
			&data.Lastname,
			&data.Email,
			&data.Phone)
		if err != nil {
			return err
		}
		member = append(member, data)
	}
	defer result.Close()
	defer db.Close()
	return c.JSON(http.StatusOK, member)
}
func updateMember(c echo.Context) error {
	id := c.Param("id")
	mb := new(Member)
	if err := c.Bind(mb); err != nil {
		return err
	}
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	sqlStatment := `UPDATE member SET username =$1, password =$2, fisrtname=$3, lastname=$4, email=$5, phone=$6 WHERE id =$7;`
	_, err = db.Exec(sqlStatment, mb.Username, mb.Password, mb.Fisrtname, mb.Lastname, mb.Email, mb.Phone, id)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "UpdateMemberSusces")
}
func deleteMember(c echo.Context) error {
	id := c.Param("id")

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	sqlStatment := `DELETE FROM member WHERE id=$1;`
	_, err = db.Exec(sqlStatment, id)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	return c.JSON(http.StatusOK, "deleteuser")
}
func login(c echo.Context) error {
	mb := new(Member)
	if err := c.Bind(mb); err != nil {
		return err
	}
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	var member Member

	sqlStatment := `SELECT id, username, password FROM member WHERE username=$1 AND password=$2;`
	err = db.QueryRow(sqlStatment, mb.Username, mb.Password).Scan(&member.ID, &member.Username, &member.Password)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, "Login Failed!!")
	}

	if member.Username == mb.Username && member.Password == mb.Password {
		fmt.Println("Login succesful!")
		return c.JSON(http.StatusOK, member)
	} else {
		fmt.Println("Login failed!")
		return c.JSON(http.StatusOK, "Login Failed!!")
	}
}

func addProduct(c echo.Context) error {
	add := new(Product)
	if err := c.Bind(add); err != nil {
		return err
	}

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	dt := time.Now()

	sqlStatment := `INSERT INTO productstore (brand, price, datetimes) VALUES ($1, $2, $3);`
	_, err = db.Exec(sqlStatment, add.Brand, add.Price, dt.Format("2006-01-02 15:04:05.000"))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "addProduct")
}
func getAllProduct(c echo.Context) error {
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Succesfully connected! DB")

	sqlStatment := `SELECT * FROM productstore;`
	result, err := db.Query(sqlStatment)
	if err != nil {
		fmt.Println(err)
	}
	var products []Product

	for result.Next() {
		var data Product
		err := result.Scan(
			&data.ProductID,
			&data.Brand,
			&data.Price,
			&data.Datetimes)
		if err != nil {
			return err
		}
		products = append(products, data)
	}
	defer result.Close()
	defer db.Close()
	return c.JSON(http.StatusOK, products)
}
func deleteProduct(c echo.Context) error {
	id := c.Param("id")

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	sqlStatment := `DELETE FROM productstore WHERE id=$1;`
	_, err = db.Exec(sqlStatment, id)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	return c.JSON(http.StatusOK, "deleteProduct")
}
func getProduct(c echo.Context) error {
	id := c.Param("id")

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	var product Product

	sqlStatment := `SELECT * FROM productstore WHERE id=$1;`
	err = db.QueryRow(sqlStatment, id).Scan(&product.ProductID, &product.Brand, &product.Price, &product.Datetimes)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return c.JSON(http.StatusOK, product)
}
func updateProduct(c echo.Context) error {
	id := c.Param("id")
	add := new(Product)
	if err := c.Bind(add); err != nil {
		return err
	}
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	sqlStatment := `UPDATE productstore SET brand =$1, price=$2 WHERE id =$3;`
	_, err = db.Exec(sqlStatment, add.Brand, add.Price, add.Datetimes, id)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "updateProduct")
}

type (
	Employee struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	Member struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Fisrtname string `json:"fisrtname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	Product struct {
		ProductID int    `json:"id"`
		Brand     string `json:"brand"`
		Price     string `json:"price"`
		Datetimes string `json:"datetimes"`
	}
)

const (
	host     = "35.240.234.180"
	port     = 80
	user     = "admin"
	password = "1234"
	dbname   = "employeesdb"
)
