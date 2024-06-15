package db

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMongoDb(t *testing.T) {
	if err := godotenv.Load(os.ExpandEnv("../../.env")); err != nil {
		t.Error("err loading env ", err)
	}
	dbUri := os.Getenv("TEST_DB_URI")
	dbName := os.Getenv("TEST_DB_NAME")

	if dbName == "" || dbUri == "" {
		t.Log("Abording test because TEST_DB_URI or TEST_DB_NAME not found in .env")
		return
	}

	ctx := context.Background()

	db := NewMongoDb(dbUri, dbName)
	db.AddCollection([]string{"tests"})

	t.Cleanup(func() {
		if err := db.Drop(); err != nil {
			t.Error("err clearing tests... ", err)
		}
	})

	t.Log("Testing connection...")
	if err := db.Connect(ctx); err != nil {
		t.Error("err connecting: ", err)
	}

	type TestData struct {
		Name string `bson:"name"`
	}

	tstD := TestData{
		Name: "test 1",
	}

	t.Log("Testing put")
	if err := db.Add("tests", tstD); err != nil {
		t.Error(err)
	}

	t.Log("Testing get")
	tstR := new(TestData)
	err := db.Get("tests", tstD, &tstR)
	if err != nil {
		t.Error(err)
	} else {
		if tstR.Name != tstD.Name {
			t.Errorf("have %s want %s", tstR.Name, tstD.Name)
		}
	}

	t.Log("Testing update")
	tstU := TestData{
		Name: "test 2",
	}
	if err := db.Update("tests", tstD, tstU); err != nil {
		t.Error(err)
	}

	t.Log("Testing deletion")
	if err := db.Delete("tests", tstD); err != nil {
		t.Error(err)
	}
	if err := db.Delete("tests", tstU); err != nil {
		t.Error(err)
	}

}
