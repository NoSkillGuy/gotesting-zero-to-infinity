package db_test

import (
	"testing"

	"github.com/NoSkillGuy/gotesting-zero-to-infinity/structuring-tests-in-go/db"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) (*db.UsersDB, func()) {
    // Initialize the DB
    database, err := db.InitializeDB()
    if err != nil {
        t.Fatalf("Failed to initialize database: %v", err)
    }

    // Create the users table
    err = database.CreateUsersTable()
    if err != nil {
        t.Fatalf("Failed to create users table: %v", err)
    }

    // Load fixture data
    err = loadFixtures(database)
    if err != nil {
        t.Fatalf("Failed to load fixtures: %v", err)
    }

   return database, func() {
        database.Db.Close()
    }
}

func loadFixtures(database *db.UsersDB) error {
    database.InsertUser("Alice", 30)
    database.InsertUser("Bob", 25)
    database.InsertUser("Charlie", 40)

    return nil
}

func TestGetUserByID(t *testing.T) {
    // Setup database and load fixtures
    database, cleanup := setupTestDB(t)
    defer cleanup()

    // Test case: Retrieve user with ID 1 (Alice)
    user, err := database.GetUserByID(1)
    if err != nil {
        t.Fatalf("Failed to get user: %v", err)
    }

    if user.Name != "Alice" || user.Age != 30 {
        t.Errorf("Expected Alice, 30; got %s, %d", user.Name, user.Age)
    }

    // Test case: Retrieve user with ID 2 (Bob)
    user, err = database.GetUserByID(2)
    if err != nil {
        t.Fatalf("Failed to get user: %v", err)
    }

    if user.Name != "Bob" || user.Age != 25 {
        t.Errorf("Expected Bob, 25; got %s, %d", user.Name, user.Age)
    }

    // Test case: User not found
    _, err = database.GetUserByID(99)
    if err == nil {
        t.Errorf("Expected error for non-existent user, got nil")
    }
}