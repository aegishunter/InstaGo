package repository

import (
	"instago/model"
	testing "testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetCountTotalPosts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Err = %s", err)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT COUNT").WillReturnRows(sqlmock.NewRows([]string{"Count"}).AddRow(1))
}

func TestPostRepository_GetPostDetailsById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s has occured", err)
	}
	defer db.Close()
	var (
		paramId     = 1
		paramUserId = 1
		rows        = []string{
			"id",
			"userId",
			"likes",
			"description",
			"createdAt",
			"updatedAt",
			"imageName",
		}
		post = model.NewPost(paramId, paramUserId, 1, "test", time.Now(), time.Now(), "test.jpg")
	)
	mock.ExpectQuery("SELECT (.+) FROM posts WHERE id ").
		WithArgs(paramId, paramUserId).
		WillReturnRows(sqlmock.NewRows(rows).
			AddRow(post.Id, post.UserId, post.Likes,
				post.Description, post.CreatedAt, post.UpdatedAt, post.ImageName))
	sql := "SELECT * FROM posts WHERE id = $1 AND \"userId\" = $2 LIMIT 1"
	data, err := db.Query(sql, paramId, paramUserId)
	if err != nil {
		t.Fatalf("Error = %s", err)
	}
	defer data.Close()
	assert.NotNil(t, data)
	assert.Nil(t, err)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("err = %s", err)
	}

}

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
