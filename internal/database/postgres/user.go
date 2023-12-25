package postgres

import (
	"context"
	"fmt"

	"github.com/PandaGoL/api-project/internal/database/postgres/models"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
)

func (pgs *PgStorage) AddOrUpdateUser(user models.User) (scanUser *models.User, err error) {

	// 0. Подготовительные операции
	if user.UserID == "" {
		user.UserID = uuid.NewV4().String()
	}

	template := `INSERT INTO myproject.users
	(%s)
	VALUES
	(%s)
	ON CONFLICT (user_id) DO UPDATE SET
	email = %s, phone = %s
	RETURNING %s`

	sql := fmt.Sprintf(template,
		user.GetColumns(),
		user.GetPlaceholders(),
		user.GetPlaceholder("email"), user.GetPlaceholder("phone"),
		user.GetColumns(),
	)

	// 1. Контекст для завершения
	ctx, cancel := context.WithTimeout(context.Background(), pgs.options.QueryTimeout)
	defer cancel()

	// 2. Запрос
	row := pgs.pool.QueryRow(ctx, sql, user.GetFieldValue()...)
	scanUser = new(models.User)
	err = row.Scan(scanUser.GetFields()...)

	//3. Статус
	switch err {
	case nil:
		return scanUser, nil
	case pgx.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (pgs *PgStorage) GetUsers() (users []*models.User, count int, err error) {
	// 0. Подготовительные операции
	//bt := time.Now()
	// 1. Подготовка запроса
	// 1.1 Базовый шаблон>
	template := `SELECT %s FROM myproject.users`
	sql := fmt.Sprintf(template, new(models.User).GetColumns())

	// 2. Контекст для завершения
	ctx1, cancel1 := context.WithTimeout(context.Background(), pgs.options.QueryTimeout)
	defer cancel1()

	// 3. Запрос
	rows, err := pgs.pool.Query(ctx1, sql)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// 4. Обработка ответа
	users = make([]*models.User, 0)
	for rows.Next() {
		var scanUser = new(models.User)
		err = rows.Scan(scanUser.GetFields()...)
		if err != nil {
			break
		}
		users = append(users, scanUser)
	}

	// 5. Получение количества записей
	sql = `SELECT COUNT(user_id) FROM myproject.users`
	// 5.1 Контекст для завершения
	ctx2, cancel2 := context.WithTimeout(context.Background(), pgs.options.QueryTimeout)
	defer cancel2()
	// 5.2 Запрос
	row := pgs.pool.QueryRow(ctx2, sql)
	// 5.3 Обработка ответа
	err = row.Scan(&count)
	switch err {
	case nil:
		return
	default:
		return nil, 0, err
	}
}

func (pgs *PgStorage) GetUser(userID string) (user *models.User, err error) {
	// 0. Подготовительные операции
	//bt := time.Now()

	// 1. Подготовка запроса
	template := `SELECT %s FROM myproject.users WHERE user_id = $1`
	sql := fmt.Sprintf(template, user.GetColumns())

	// 2. Контекст для завершения
	ctx, cancel := context.WithTimeout(context.Background(), pgs.options.QueryTimeout)
	defer cancel()

	// 3. Запрос
	row := pgs.pool.QueryRow(ctx, sql, userID)

	// 4. Обработка ответа
	var scanPartner = new(models.User)
	err = row.Scan(scanPartner.GetFields()...)
	switch err {
	case nil:
		//TODO метрики
		return scanPartner, nil
	case pgx.ErrNoRows:
		//TODO метрики
		return nil, nil
	default:
		//TODO метрики
		return nil, err
	}
}

func (pgs *PgStorage) DeleteUser(userID string) error {
	// 0. Подготовительные операции
	//bt := time.Now()

	// 1. Подготовка запроса
	template := `DELETE FROM myproject.users WHERE user_id = $1`

	// 2. Контекст для завершения
	ctx, cancel := context.WithTimeout(context.Background(), pgs.options.QueryTimeout)
	defer cancel()

	// 3. Запрос
	_, err := pgs.pool.Exec(ctx, template, userID)

	//4. Статус
	switch err {
	case nil:
		//TODO метрики
		return nil
	case pgx.ErrNoRows:
		//TODO метрики
		return nil
	default:
		//TODO метрики
		return err
	}
}
