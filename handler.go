package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	demo "rpc_server/kitex_gen/demo"
	"strconv"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

const (
	host     = "corax.com.cn"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "student"
)

func QueryFromDatabase(id int32, student *demo.Student) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return err
	}
	student.Id = -1

	query, err := database.Query("select * from student where id = " + strconv.Itoa(int(id)))
	if err != nil {
		log.Fatal(err)
		return err
	}

	var cid int
	student.College = new(demo.College)
	for query.Next() {
		err = query.Scan(&student.Id, &student.Name, &cid, &student.Sex)
	}

	if student.Id == -1 {
		return nil
	}

	query, err = database.Query("select * from college where id = " + strconv.Itoa(cid))
	for query.Next() {
		err = query.Scan(&cid, &student.College.Name, &student.College.Address)
	}

	query, err = database.Query("select * from email where id = " + strconv.Itoa(int(id)))

	var emails []string

	for query.Next() {
		var email string
		err = query.Scan(&student.Id, &email)
		emails = append(emails, email)
	}

	student.Email = emails

	err = query.Close()

	err = database.Close()

	return nil
}

func InsertIntoDatabase(student *demo.Student) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return err
	}

	tx, err := database.Begin()

	var cid = -1
	var cname string
	var caddr string
	query, err := database.Query("select * from college where name = " + "'" + student.College.Name + "'")
	for query.Next() {
		err = query.Scan(&cid, &cname, &caddr)
	}
	if cid == -1 {
		stmt, _ := tx.Prepare("insert into college(name, address) values ($1, $2)")
		_, err = stmt.Exec(student.College.Name, student.College.Address)
		err = tx.Commit()

		query, err = database.Query("select * from college where name = " + "'" + student.College.Name + "'")
		for query.Next() {
			err = query.Scan(&cid, &cname, &caddr)
		}

		tx, err = database.Begin()
	}
	stmt, err := tx.Prepare("insert into student values ($1, $2, $3, $4)")

	_, err = stmt.Exec(student.Id, student.Name, cid, student.Sex)

	for i := 0; i < len(student.Email); i++ {
		stmt, err = tx.Prepare("insert into email values ($1, $2)")
		_, err = stmt.Exec(student.Id, student.Email[i])
	}

	err = tx.Commit()
	return nil
}

var studentMap = make(map[int32]*demo.Student)

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	resp = demo.NewRegisterResp()
	var newStudent demo.Student
	err = QueryFromDatabase(student.Id, &newStudent)
	if err != nil {
		resp.Success = false
		resp.Message = "Internal Exception"
	}
	if newStudent.Id > 0 {
		resp.Success = false
		resp.Message = "Register Failed: Student Information Already Exists"
	} else {
		err = InsertIntoDatabase(student)
		if err != nil {
			resp.Success = false
			resp.Message = "Internal Exception"
		}
		resp.Success = true
		resp.Message = "Register Success"
	}
	fmt.Println(resp)
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	resp = demo.NewStudent()
	var oldStudent demo.Student
	if value, exist := studentMap[req.Id]; exist {
		fmt.Println("Use Cache")
		resp = value
		return
	} else {
		fmt.Println("Query Database")
		err = QueryFromDatabase(req.Id, &oldStudent)
		if err != nil {
			return
		}
		if oldStudent.Id == -1 {
			var student = demo.Student{
				Id:      -1,
				Name:    "Student Not Exist",
				College: &demo.College{Name: "Unknown", Address: "Unknown"},
				Email:   nil,
			}
			resp = &student
		} else {
			resp = &oldStudent
			studentMap[req.Id] = &oldStudent
		}
		return
	}
}

// GetPort implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) GetPort(ctx context.Context, req *demo.GetPortReq) (resp *demo.GetPortResp, err error) {
	resp = new(demo.GetPortResp)
	resp.Port = strconv.Itoa(bindPort)
	return
}
