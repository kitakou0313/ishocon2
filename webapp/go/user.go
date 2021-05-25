package main

import "errors"

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
	myNumberToUserCache = map[string]User{}

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
		key := u.Name + u.Address + u.MyNumber
		myNumberToUserCache[key] = u
	}

}

func getUser(name string, address string, myNumber string) (User, error) {
	key := name + address + myNumber
	if _, ok := myNumberToUserCache[key]; !ok {
		return User{}, errors.New("There is not user with mynumber")
	}
	return myNumberToUserCache[key], nil
}
