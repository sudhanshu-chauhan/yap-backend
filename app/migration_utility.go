package app

var models []interface{} = []interface{}{
	User{},
	Task{},
}

func MigrateTables() {
	dbconn := GetConnection()
	for _, currentModel := range models {
		if !(dbconn.Migrator().HasTable(&currentModel)) {
			dbconn.Migrator().CreateTable(&currentModel)
		}
	}

}
