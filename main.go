package main

// ***REMEMBER TO UNCOMMENT IMPORT***

// import (
//   "fmt"
//   "net/http" //http request handle karne ke liye

//   "github.com/labstack/echo/v4" //Echo web framework ke liye
//   "./user" //user.go file ko import karne ke liye
// )

// In-memory user storage (database se replace kardena baadmein)
var users []user.User

func main() {
	//Echo instance taiyyar karna hai
	e := echo.New()

	//user registration ke lilye route
	e.POST("/register", user.RegisterHandler)

	//server start karte hain
	e.Logger.Fatal(e.Start(":8080")) // baad mein desired port number se replace karna hai
}

//abhi ke liye itna rehne dete hain.. baad ka dekhta hoon; pehle thoda user.go ka maamla dekhleta
