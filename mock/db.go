package main

type DB interface {
    Get(key string) (int, error)
}

func getFromDB(db DB, key string) int {
    if value, err = db.Get(key); err != nil {
        return value
    }
    return -1
}
