package tamiat_user

import (
	"database/sql"
	"errors"
	"github.com/tamiat/backend/pkg/errs"
	"gorm.io/gorm"
)

type TamiatUserRepositoryDb struct {
	db  *sql.DB
	gDb *gorm.DB
}

func (r TamiatUserRepositoryDb) Login(userObj TamiatUser) (string, error) {
	var retrievedUsr TamiatUser
	if err := r.gDb.Where("email = ?", userObj.Email).First(&retrievedUsr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errs.ErrRecordNotFound
		}
		return "", errs.ErrDb
	}
	return retrievedUsr.Password, nil
}
func (r TamiatUserRepositoryDb) Create(tUserObj TamiatUser) (int, error) {
	var id int
	row := r.db.QueryRow(`INSERT INTO tamiat_users(name,email,password,role_id) VALUES ($1,$2,$3,$4) RETURNING id`, tUserObj.Name, tUserObj.Email, tUserObj.Password, tUserObj.RoleId)
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (r TamiatUserRepositoryDb) ReadAll() ([]TamiatUser, error) {
	var tamiatUser TamiatUser
	var tamiatUserArr []TamiatUser
	rows, err := r.db.Query(`SELECT id,created_at,updated_at,deleted_at,name,email,role_id FROM tamiat_users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&tamiatUser.Id, &tamiatUser.CreatedAt, &tamiatUser.UpdatedAt, &tamiatUser.DeletedAt, &tamiatUser.Name, &tamiatUser.Email, &tamiatUser.RoleId)
		if err != nil {
			return nil, err
		}
		tamiatUserArr = append(tamiatUserArr, tamiatUser)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tamiatUserArr, nil
}
func (r TamiatUserRepositoryDb) ReadUserByID(id int) (TamiatUser, error) {
	row := r.db.QueryRow(`SELECT id,created_at,updated_at,deleted_at,name,email,role_id FROM tamiat_users WHERE id=$1`, id)
	var tamiatUser TamiatUser
	err := row.Scan(&tamiatUser.Id, &tamiatUser.CreatedAt, &tamiatUser.UpdatedAt, &tamiatUser.DeletedAt, &tamiatUser.Name, &tamiatUser.Email, &tamiatUser.RoleId)
	if err == sql.ErrNoRows {
		return tamiatUser, sql.ErrNoRows
	}
	return tamiatUser, err
}
func (r TamiatUserRepositoryDb) Update(tUserObj TamiatUser) error {
	row := r.db.QueryRow(`UPDATE tamiat_users SET name = $1,role_id = $2 WHERE id = $3 RETURNING id;`, tUserObj.Name, tUserObj.RoleId, tUserObj.Id)
	var id int
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return sql.ErrNoRows
	}
	return err
}
func (r TamiatUserRepositoryDb) ResetPassword(tUserObj TamiatUser) error {
	row := r.db.QueryRow(`UPDATE tamiat_users SET password = $1 WHERE id = $2 RETURNING id;`, tUserObj.Password, tUserObj.Id)
	var id int
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return sql.ErrNoRows
	}
	return err
}
func (r TamiatUserRepositoryDb) Delete(id int) error {
	row := r.db.QueryRow(`DELETE FROM tamiat_users WHERE id = $1 RETURNING id;`, id)
	var tmpId int
	err := row.Scan(&tmpId)
	if err == sql.ErrNoRows {
		return sql.ErrNoRows
	}
	return err
}
func (r TamiatUserRepositoryDb) GetRoleId(name string) (int, error) {
	var id int
	row := r.db.QueryRow(`SELECT id FROM tamiat_roles WHERE name=$1`, name)
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
func NewTamiatUserRepositoryDb(db *sql.DB, gDb *gorm.DB) TamiatUserRepositoryDb {
	return TamiatUserRepositoryDb{db, gDb}
}
