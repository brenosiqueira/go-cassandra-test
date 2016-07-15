package main

import (
  "fmt"
  "log"
  "github.com/gocql/gocql"


)

func main() {
  // connect to the cluster 10.0.75.1
  cluster := gocql.NewCluster("localhost")
	//cluster := gocql.NewCluster("10.0.75.1")

	cluster.Keyspace = "dev"
  cluster.ProtoVersion = 3 // muda conforme a versão do cassandra
	session, _ := cluster.CreateSession()

  // insert a user
  if err := session.Query("insert into emp (empid, emp_first, emp_last, emp_dept) values (3,'fred','smith','eng')").Exec(); err != nil {
  		log.Fatal(err)
  }

  // Use select to get the user we just entered
	var emp_dept, emp_first, emp_last string
	var empid int

	if err2 := session.Query("select empid, emp_dept, emp_first, emp_last from emp where  empid = 1").Scan(&empid, &emp_dept, &emp_first, &emp_last); err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(empid, emp_dept, emp_first, emp_last)

  defer session.Close()
}
