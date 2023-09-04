package database

import (
	"database/sql"
	m "tg-group-scheduler/models"
)

type Database interface {
	FindUserById(userId int) (*m.User, error)
	FindUserByTelegramId(telegramId int) (*m.User, error)
	InsertUser(user m.User) (*m.User, error)
}

type PostgresDatabase struct {
	connectionString string
	db               sql.DB
}

func (p *PostgresDatabase) FindUserById(userId int) (*m.User, error) {
	sql := `
      select * from $1
      where id = $2
    `

	row := p.db.QueryRow(sql, m.USER_TABLE, userId)
	var user = m.User{}

	err := row.Scan(user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PostgresDatabase) FindUserByTelegramId(telegramId int) (*m.User, error) {
	sql := `
    select * from $1
    where telegram_id = $2
  `

	row := p.db.QueryRow(sql, m.USER_TABLE, telegramId)
	var user = m.User{}

	err := row.Scan(user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PostgresDatabase) InsertUser(user m.User) (*m.User, error) {
	sql := `
    insert into $1
    values (telegram_id, default_timezone)
  `
	res, err := p.db.Exec(sql, user.TelegramId, user.DefaultTimezone)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	newUser := m.NewUser(
		int(id),
		user.TelegramId,
		user.DefaultTimezone,
	)

	return newUser, err
}
