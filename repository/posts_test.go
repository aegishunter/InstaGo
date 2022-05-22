package repository

import (
	testing "testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPostRepository_UpdateDescription(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s has occured", err)
	}
	defer db.Close()
	var (
		paramDescription = "Test Hello"
		paramId          = 1
	)
	mock.ExpectExec("UPDATE posts SET description").
		WithArgs(paramDescription, paramId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	sql := "UPDATE posts SET description = $1, \"updatedAt\" = now() WHERE id = $2"
	_, err = db.Exec(sql, paramDescription, paramId)
	if err != nil {
		t.Fatalf("err = %s", err)
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("err mock expectations = %s", err)
	}
}

func TestPostRepository_AddPost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s has occured", err)
	}
	defer db.Close()

	var (
		paramUserId      = 1
		paramDescription = "test"
		paramImageName   = "test.jpg"
	)

	mock.ExpectExec("INSERT INTO posts").
		WithArgs(paramUserId, paramDescription, paramImageName).
		WillReturnResult(sqlmock.NewResult(1, 1))

	sql := "INSERT INTO posts(\"userId\", likes, description, \"imageName\") values ($1,0,$2,$3)"
	_, err = db.Exec(sql, paramUserId, paramDescription, paramImageName)
	if err != nil {
		t.Fatalf("an error %s has occured", err)
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("expectation not met. err = %s", err)
	}
}
