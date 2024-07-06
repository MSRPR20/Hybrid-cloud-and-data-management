package database

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func ConnectSQLite(dbPath string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }

    query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );
    `
    _, err = db.Exec(query)
    if err != nil {
        return nil, err
    }

    return db, nil
}

func GetUsers(db *sql.DB) ([]map[string]interface{}, error) {
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []map[string]interface{}
    for rows.Next() {
        var id int
        var name, email string
        if err := rows.Scan(&id, &name, &email); err != nil {
            return nil, err
        }
        users = append(users, map[string]interface{}{
            "id":    id,
            "name":  name,
            "email": email,
        })
    }

    return users, nil
}
