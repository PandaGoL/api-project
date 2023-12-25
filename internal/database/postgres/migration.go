package postgres

import (
	"database/sql"
	"os"

	"github.com/PandaGoL/api-project/internal/database/postgres/types"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

func (pgs *PgStorage) Migrations() error {
	// 0. Подготвительные операции
	goose.SetTableName("db_version")
	logrus.Warn("Start process migration")
	if !pgs.options.MigrationEnable {
		return types.ErrMigrationsNotEnable
	}

	// 1. Парсинг конфига и заполнение нужными данными

	conf, err := pgx.ParseConfig(pgs.options.DSN())
	if err != nil {
		return err
	}

	connStr := stdlib.RegisterConnConfig(conf)
	// 2. Создание соединения
	var dbConn *sql.DB
	dbConn, err = sql.Open("pgx", connStr)
	if err != nil {
		logrus.WithError(err).Error("Unable to make connection")
		return err
	} else {
		logrus.Warn("Connection was established")
	}
	defer func() {
		err = dbConn.Close()
		if err != nil {
			logrus.WithError(err).Error("Unable to close postgres-connection")
		}
	}()

	// 3. Проверка текущей версии
	version, err := goose.EnsureDBVersion(dbConn)
	if err != nil {
		logrus.WithError(err).Error("Unable obtain current version")
		return err
	} else {
		logrus.WithField("version", version).Warn("Current Version migrations")
	}

	// 4. Открытие текущей дирректории с миграциями
	dir, _ := os.Getwd()
	migrations, err := os.ReadDir("./internal/migrations")
	var fileNames []string
	for _, file := range migrations {
		fileNames = append(fileNames, file.Name())
	}
	if err != nil {
		logrus.WithError(err).Error("Unable read current directory")
		return err
	} else {
		logrus.WithField("dir", dir).WithField("/internal/migrations ->", fileNames).Debug("Current directory")
	}

	// 5. Установка актуальной версии миграции
	err = goose.UpTo(dbConn, "./internal/migrations", pgs.options.MigrationVersion)
	if err != nil {
		logrus.WithError(err).Error("Unable set current version migrations")
		return err
	} else {
		logrus.Debug("Current version migrations set successfully")
	}

	// 6. Проверка текущей версии
	version, err = goose.EnsureDBVersion(dbConn)
	if err != nil {
		logrus.WithError(err).Error("Unable check current version migration")
		return err
	} else {
		logrus.WithField("current_version", version).Debug("Updated to migration")
	}

	return nil

}
