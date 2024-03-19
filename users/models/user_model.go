package models

type User struct {
    ID        int    `db:"id"`
    CreatedAt string `db:"created_at"`
    UpdatedAt string `db:"updated_at"`
    Name      string `db:"name"`
    Email     string `db:"email"`
    Phone     string `db:"phone"`
}
