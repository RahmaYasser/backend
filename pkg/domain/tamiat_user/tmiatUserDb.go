package tamiat_user

import (
	"database/sql"
)

type TamiatUserRepositoryDb struct {
	db *sql.DB
}

func (r TamiatUserRepositoryDb) Create(tUserObj TamiatUser) error {
	_, err := r.db.Exec(`INSERT INTO tamiat_users(name,email,password,role_id) VALUES ($1,$2,$3,$4)`, tUserObj.Name, tUserObj.Email, tUserObj.Password, tUserObj.RoleId)
	return err
}

func (r TamiatUserRepositoryDb) ReadAll() ([]TamiatUser, error) {
	var tamiatUser TamiatUser
	var tamiatUserArr []TamiatUser
	rows, err := r.db.Query(`SELECT * FROM tamiat_users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&tamiatUser)
		tamiatUserArr = append(tamiatUserArr, tamiatUser)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tamiatUserArr, nil
}
func NewTamiatUserRepositoryDb(db *sql.DB) TamiatUserRepositoryDb {
	return TamiatUserRepositoryDb{db}
}
