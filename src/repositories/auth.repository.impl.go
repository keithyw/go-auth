package repositories

import (
	"github.com/keithyw/auth/database"
	"github.com/keithyw/auth/models"
)

type AuthRepositoryImpl struct {
	Conn *database.MysqlDB
}

func NewAuthRepository(conn *database.MysqlDB) AuthRepository {
	return &AuthRepositoryImpl{conn}
}

func (a *AuthRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	var user models.User
	stmt, err := a.Conn.DB.Prepare("SELECT * FROM users WHERE username = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.Passwd)
	if err != nil {
		return nil, err
	}
	return &user, nil
}