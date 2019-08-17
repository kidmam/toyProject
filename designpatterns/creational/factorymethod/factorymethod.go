package factorymethod

import "fmt"

type (
	Database interface {
		GetData(string) string
		PutDate(string, string)
	}

	mongoDB struct {
		database map[string]string
	}

	sqlite struct {
		database map[string]string
	}
)

func (mdb mongoDB) GetData(query string) string {
	fmt.Printf("MongoDB")
	return mdb.database[query]
}

func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sq sqlite) GetData(query string) string {
	fmt.Printf("Sqlite")
	return sq.database[query]
}

func (sq sqlite) PutData(query string, data string) {
	sq.database[query] = data
}

func DatabaseFactory(env string) Database {
	switch env {
	case "production":
		return mongoDB {
			database: make(map[string]string)
		}
	case "development":
		return sqlite {
			database: make(map[string]string)
		}
	default:
		return nil
	}
}