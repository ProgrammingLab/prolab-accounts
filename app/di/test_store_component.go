package di

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

// MustCreateTestStoreComponent creates test store component or exits
func MustCreateTestStoreComponent(cfg *config.Config) *TestStoreComponent {
	db := mustConnectTestRDB(cfg)

	boil.SetDB(db)

	return &TestStoreComponent{
		storeComponentImpl: &storeComponentImpl{
			db:  sqlutil.New(db),
			cfg: cfg,
		},
	}
}

func mustConnectTestRDB(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.TestDataBaseURL)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("%+v", err)
	}

	return db
}

type TestStoreComponent struct {
	*storeComponentImpl
}

func (c *TestStoreComponent) MustClose() {
	t := record.TableNames
	names := []string{
		t.Blogs,
		t.Departments,
		t.Entries,
		t.Profiles,
		t.Roles,
		t.Users,
	}
	q := fmt.Sprintf("truncate %s cascade", strings.Join(names, ", "))
	_, err := c.db.Exec(q)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	err = c.db.Close()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
