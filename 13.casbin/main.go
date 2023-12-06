package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")

	// Load policies from database
	a, _ := gormadapter.NewAdapter("mysql", "go:1234@tcp(127.0.0.1:3306)/go", true)
	e, _ := casbin.NewEnforcer("./model.conf", a)

	// Simulate request
	sub := "alice" // The user that wants to access a resource.
	obj := "data1" // The resource that is going to be accessed.
	act := "read"  // The operation that the user performs on the resource.

	// Add policy
	// added, err := e.AddPolicy("alice", "data1", "read")
	// added, err = e.AddPolicy("alice", "data2", "read")
	// fmt.Println(added, '\n', err)

	// Select
	// filteredPolicy := e.GetFilteredPolicy(1, "data1")
	// fmt.Println(filteredPolicy)

	// Update
	// updated, err := e.UpdatePolicy([]string{"alice", "data1", "read"}, []string{"alice", "data3", "write"})
	// fmt.Println(updated, err)

	added, err := e.AddGroupingPolicy("alice", "data2_admin")
	fmt.Println(added, err)

	// Add function
	e.AddFunction("my_func", KeyMatchFunc)

	// Checking
	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// Handle err
		fmt.Printf("%s", err)
	}

	if ok {
		// Permit alice to read data1
		fmt.Println("Passed!")
	} else {
		// Denied request, throw exception
		fmt.Println("Not pass.")
	}

}

// Customize function
func KeyMatch(key1 string, key2 string) bool {
	return true
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}
