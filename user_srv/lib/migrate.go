package lib

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrate struct {
	conn  *migrate.Migrate
	dbConn *sql.DB
	migrateLink string 
	migratePath string
	databaseName string
}

func NewMigrate(databaseName string, migrateLink string, migratePath string) *Migrate {
	return &Migrate{
    migrateLink: migrateLink,
		migratePath: migratePath,
		databaseName: databaseName,
	}
}

func (m *Migrate) getConn () error {
	db, err := sql.Open("mysql", m.migrateLink)

	if err != nil {
		log.Printf("[ERROR] lib.migrate.getConn.Open error: %v", err)
		return err
	}

  // instance must have `multiStatements` set to true
	driver, err := mysql.WithInstance(db, &mysql.Config{})	


	if err != nil {
		log.Printf("[ERROR] lib.migrate.getConn.WithInstance error: %v", err)
		return err
	}

	migratePath := fmt.Sprintf("file://%s", m.migratePath)
	migrateInstance, err := migrate.NewWithDatabaseInstance(
		  migratePath,
			m.databaseName, 
			driver,
	)
	if err != nil {
		log.Printf("[ERROR] lib.migrate.getConn.NewWithDatabaseInstance error: %v", err)
		return err
	}

	m.SetConn(migrateInstance)
	m.dbConn = db
	return nil
}

func (m *Migrate) SetConn(conn *migrate.Migrate) {
	m.conn = conn
}

func (m *Migrate) MigrateUpDatabase() error {
	defer func () {
		if err := m.dbConn.Close(); err != nil {
      log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Close error: %v", err)
		} else {
			log.Println("[INFO] migrate sql conn close successfully~~")
		}
	}()

  m.getConn()
	err := m.conn.Up()


	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Up error: %v", err)
		return nil
	}

	if err != nil {
		log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Up error: %v", err)
		return err
	}

	log.Println("[INFO] migrate successfully up~~")
	return nil
}

func (m *Migrate) MigrateStepDatabase(n int) error {
	defer func () {
		if err := m.dbConn.Close(); err != nil {
      log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Close error: %v", err)
		} else {
			log.Println("[INFO] migrate sql conn close successfully~~")
		}
	}()

  m.getConn()
	err := m.conn.Steps(n)

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Up error: %v", err)
		return nil
	}

	if err != nil {
		log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Up error: %v", err)
		return err
	}

	log.Println("[INFO] migrate successfully step~~")
	return nil
}

func (m *Migrate) MigrateForceDatabase(n int) error {
	defer func () {
		if err := m.dbConn.Close(); err != nil {
      log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Close error: %v", err)
		} else {
			log.Println("[INFO] migrate sql conn close successfully~~")
		}
	}()
	
  m.getConn()
	err := m.conn.Force(n)
	if err != nil {
		log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Up error: %v", err)
		return err
	}

	log.Println("[INFO] migrate successfully force~~")
	return nil
}

func (m *Migrate) MigrateDownDatabase() error {
	defer func () {
		if err := m.dbConn.Close(); err != nil {
      log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Close error: %v", err)
		} else {
			log.Println("[INFO] migrate sql conn close successfully~~")
		}
	}()

  m.getConn()
	err := m.conn.Down()


	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Up error: %v", err)
		return nil
	}

	if err != nil {
		log.Printf("[ERROR] lib.migrate.getConn.MigrateUpDatabase.Up error: %v", err)
		return err
	}

	log.Println("[INFO] migrate successfully down~~")
	
	return nil
}