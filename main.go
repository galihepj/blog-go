package main

import (
	"fmt"
	"database/sql"  // Database SQL package to perform queries
	"log"           // Display messages to console
	"net/http"      // Manage URL
	"text/template" // Manage HTML files

	_ "github.com/go-sql-driver/mysql" // MySQL Database driver
)

  var logn int = 0
  var emailnya string ="admin@gmail.com"
  var passnya string ="admin"

// Struct used to send data to template
// this struct is the same as the database
type Names struct {
	Id    int
	Name  string
	Email string
}

// Function dbConn opens connection with MySQL driver
// send the parameter `db *sql.DB` to be used by another functions
func dbConn() (db *sql.DB) {

	dbDriver := "mysql"   // Database driver
	dbUser := "root"      // Mysql username
	dbPass := "" // Mysql password
	dbName := "db"   // Mysql schema

	// Realize the connection with mysql driver
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// If error stop the application
	if err != nil {
		panic(err.Error())
	}

	// Return db object to be used by another functions
	return db
}


// Read all templates on folder `tmpl/*`
var tmpl = template.Must(template.ParseGlob("tmpl/*"))

/////// Function Fiew shows all values on home
func Index(w http.ResponseWriter, r *http.Request) {
	
	// (View the file: `tmpl/Index`
	tmpl.ExecuteTemplate(w, "Index", nil)

}

///////22 Function Fiew shows all values on home
func Login(w http.ResponseWriter, r *http.Request) {
	
	// (View the file: `tmpl/Index`
	tmpl.ExecuteTemplate(w, "Login", nil)

}

///////22 Function Fiew shows all values on home
func Logout(w http.ResponseWriter, r *http.Request) {
	
	logn=0
	
	// (View the file: `tmpl/Index`
	tmpl.ExecuteTemplate(w, "Index", nil)

}
///////33 Function Fiew shows all values on home
func Indexadmin(w http.ResponseWriter, r *http.Request) {
	
	// (View the file: `tmpl/Index`
	tmpl.ExecuteTemplate(w, "IndexAdmin", nil)

}


// Function Fiew shows all values on home
func View(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	// Prepare a SQL query to select all data from database and threat errors
	selDB, err := db.Query("SELECT * FROM names ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := Names{}

	// Create a slice to store all data from struct
	res := []Names{}

	// Read all rows from database
	for selDB.Next() {
		// Must create this variables to store temporary query
		var id int
		var name, email string

		// Scan each row storing values from the variables above and check for errors
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.Id = id
		n.Name = name
		n.Email = email

		// Join each row on struct inside the Slice
		res = append(res, n)

	}

	// Execute template `Index` from `tmpl/*` folder and send the struct
	// (View the file: `tmpl/Index`
	tmpl.ExecuteTemplate(w, "View", res)

	// Close database connection
	defer db.Close()
}

// Function Show displays a single value
func Show(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	// Get the URL `?id=X` parameter
	nId := r.URL.Query().Get("id")

	// Perform a SELECT query getting the register Id(See above) and check for errors
	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := Names{}

	// Read all rows from database
	// This time we are going to get only one value, doesn't need the slice
	for selDB.Next() {
		// Store query values on this temporary variables
		var id int
		var name, email string

		// Scan each row to match the ID and check for errors
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.Id = id
		n.Name = name
		n.Email = email

	}
	
	
	
	// Execute template `Show` from `tmpl/*` folder and send the struct
	// (View the file: `tmpl/Show`)
	tmpl.ExecuteTemplate(w, "Show", n)

	// Close database connection
	defer db.Close()

}

// Function New just parse a form to send data to Insert function
// (View the file: `tmpl/New`)
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Function Edit works like Show
// Only select the values to send to the Edit page Form
// (View the file: `tmpl/Edit`)
func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	// Get the URL `?id=X` parameter
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	n := Names{}

	for selDB.Next() {
		var id int
		var name, email string

		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		n.Id = id
		n.Name = name
		n.Email = email

	}

	tmpl.ExecuteTemplate(w, "Edit", n)

	defer db.Close()
}

// Function Insert puts data into the database
func Insert(w http.ResponseWriter, r *http.Request) {

	// Open database connection
	db := dbConn()

	// Check the request form METHOD
	if r.Method == "POST" {

		// Get the values from Form
		name := r.FormValue("name")
		email := r.FormValue("email")

		// Prepare a SQL INSERT and check for errors
		insForm, err := db.Prepare("INSERT INTO names(name, email) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}

		// Execute the prepared SQL, getting the form fields
		insForm.Exec(name, email)

		// Show on console the action
		log.Println("INSERT: Name: " + name + " | E-mail: " + email)
	}

	// Close database connection
	defer db.Close()

	// Redirect to HOME
	http.Redirect(w, r, "/viewadmin", 301)
}

// Function Update, update values from database,
// It's the same as Insert and New
func Update(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	if r.Method == "POST" {

		// Get the values from form
		name := r.FormValue("name")
		email := r.FormValue("email")
		id := r.FormValue("uid") // This line is a hidden field on form (View the file: `tmpl/Edit`)

		// Prepare the SQL Update
		insForm, err := db.Prepare("UPDATE names SET name=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		// Update row based on hidden form field ID
		insForm.Exec(name, email, id)

		// Show on console the action
		log.Println("UPDATE: Name: " + name + " | E-mail: " + email)
	}

	defer db.Close()

	// Redirect to Home
	http.Redirect(w, r, "/viewadmin", 301)
}

// Function Delete destroys a row based on ID
func Delete(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	// Get the URL `?id=X` parameter
	nId := r.URL.Query().Get("id")

	// Prepare the SQL Delete
	delForm, err := db.Prepare("DELETE FROM names WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	// Execute the Delete SQL
	delForm.Exec(nId)

	// Show on console the action
	log.Println("DELETE")

	defer db.Close()

	// Redirect a HOME
	http.Redirect(w, r, "/viewadmin", 301)
}


func Feriv(w http.ResponseWriter, r *http.Request) {

	

	// Check the request form METHOD
	if r.Method == "POST" {

		// Get the values from Form
		email := r.FormValue("email")
		pass := r.FormValue("pass")
		
	
		if email == emailnya &&  pass == passnya {
			logn=1
			fmt.Fprintf(w, `
			<html>
			<head>
				<!-- Latest compiled and minified CSS -->
				<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">

				<!-- Optional theme -->
				<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap-theme.min.css" integrity="sha384-fLW2N01lMqjakBkx3l/M9EahuwpSfeNvV63J5ezn3uZzapT0u7EYsXMjQV+0En5r" crossorigin="anonymous">

				<!-- Latest compiled and minified JavaScript -->
				<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
			</head>
			<body>
				<div class="col-md-8 col-md-offset-2" style="margin-top:30px;">
					<div class="col-md-8 col-md-offset-2">

						<h2>Login Successful.... Welcome admin <br/>m(_ _)m</h2>
						<button type="button" class="btn btn-success"><a href="/admin" style="text-decoration:none">Home</a></button><br/>
					</div>
				</div>
			</body></html>
			`) 
		}else {
			fmt.Fprintf(w, `
			<html>
			<head>
				<!-- Latest compiled and minified CSS -->
				<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">

				<!-- Optional theme -->
				<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap-theme.min.css" integrity="sha384-fLW2N01lMqjakBkx3l/M9EahuwpSfeNvV63J5ezn3uZzapT0u7EYsXMjQV+0En5r" crossorigin="anonymous">

				<!-- Latest compiled and minified JavaScript -->
				<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
			</head>
			<body>
				<div class="col-md-8 col-md-offset-2" style="margin-top:30px;">
					<div class="col-md-8 col-md-offset-2">
						<h2>Login Failed..!!!</h2>
						Choose action :</br>
						<a href="/" >Home</a> &nbsp;&nbsp;&nbsp;
						<a href="/login" >Login Again</a>
					</div>
				</div>
			</body></html>
			`) 
		}
	
	}
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Show on console the application stated
	log.Println("Server started on: http://localhost:9000")

	// URL management
	// Manage templates
	http.HandleFunc("/", Index)    // INDEX :: Show all registers
	http.HandleFunc("/view", View) // INDEX :: Show all registers
	http.HandleFunc("/show", Show) // SHOW  :: Show only one register
	http.HandleFunc("/new", New)   // NEW   :: Form to create new register
	http.HandleFunc("/edit", Edit) // EDIT  :: Form to edit register
	
	http.HandleFunc("/login", Login) // INDEX :: Show all registers
	http.HandleFunc("/logout", Logout) // INDEX :: Show all registers
	http.HandleFunc("/admin", Indexadmin)    // INDEX :: Show all registers
	http.HandleFunc("/viewadmin", ViewAdmin) // INDEX :: Show all registers 
	http.HandleFunc("/showadmin", ShowAdmin) // INDEX :: Show all registers
	
	// Manage actions
	http.HandleFunc("/feriv", Feriv) // INSERT :: New register
	http.HandleFunc("/insert", Insert) // INSERT :: New register
	http.HandleFunc("/update", Update) // UPDATE :: Update register
	http.HandleFunc("/delete", Delete) // DELETE :: Destroy register

	// Start the server on port 9000
	http.ListenAndServe(":9000", nil)

}




///
func ViewAdmin(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	// Prepare a SQL query to select all data from database and threat errors
	selDB, err := db.Query("SELECT * FROM names ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := Names{}

	// Create a slice to store all data from struct
	res := []Names{}

	// Read all rows from database
	for selDB.Next() {
		// Must create this variables to store temporary query
		var id int
		var name, email string

		// Scan each row storing values from the variables above and check for errors
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.Id = id
		n.Name = name
		n.Email = email

		// Join each row on struct inside the Slice
		res = append(res, n)

	}

	// Execute template `Index` from `tmpl/*` folder and send the struct
	// (View the file: `tmpl/Index`
	tmpl.ExecuteTemplate(w, "ViewAdmin", res)

	// Close database connection
	defer db.Close()
}
// Function Show displays a single value
func ShowAdmin(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	// Get the URL `?id=X` parameter
	nId := r.URL.Query().Get("id")

	// Perform a SELECT query getting the register Id(See above) and check for errors
	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := Names{}

	// Read all rows from database
	// This time we are going to get only one value, doesn't need the slice
	for selDB.Next() {
		// Store query values on this temporary variables
		var id int
		var name, email string

		// Scan each row to match the ID and check for errors
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.Id = id
		n.Name = name
		n.Email = email

	}
	
	
	
	// Execute template `Show` from `tmpl/*` folder and send the struct
	// (View the file: `tmpl/Show`)
	tmpl.ExecuteTemplate(w, "ShowAdmin", n)

	// Close database connection
	defer db.Close()

}



