package main

// User Model
type User struct {
	ID       int
	Name     string
	Address  string
	MyNumber string
	Votes    int
}

var myNumberToUserCache map[string]User

func cacheAllUsers() {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Address, &u.MyNumber, &u.Votes)
		if err != nil {
			panic(err.Error())
		}

		myNumberToUserCache[u.MyNumber] = u

	}

}

func getUser(name string, address string, myNumber string) (user User, err error) {
	return myNumberToUserCache[myNumber], nil
}
