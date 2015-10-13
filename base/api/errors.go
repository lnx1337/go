package api

var (
	ErrNoSuchRow        = `ERR_NO_SUCH_ROW`             // code 1
	ErrDuplicatedRow    = `ERR_DUPLICATED_ROW`			// code 2
	ErrExpectingInteger = `ERR_EXPECTING_INTEGER`		// code 3
	ErrExpectingString  = `ERR_EXPECTING_STRING`		// code 4
	ErrUrl              = `ERR_URL`						// code 5
	ErrNotFound         = `ERR_NOT_FOUND`				// code 6
	ErrEmail            = `ERR_EMAIL`					// code 7
	ErrDbInstert        = `ERR_DB_INSERT`				// code 8
	ErrNoPrimaryKeyFound= `ERR_NO_PRIMARY_KEY_FOUND`	// code 9
	ErrSelectAll 		= `ERR_SELECT_ALL` 				// code 10
	ErrDbUpdate			= `ERR_DB_UPDATE`				// code 11
	ErrJsonBody			= `ERR_JSON_BODY`				// code 12
	ErrPasswordNoMatch	= `ERR_PASSWORD_NO_MATCH`		// code 13
	ErrLogin			= `ERR_LOGIN`					// code 14
	ErrDeleteFile		= `ERR_DELETE_FILE`				// code 15
	ErrSqlError			= `ERR_SQL_ERROR`				// code 16
	ErrMongoSelectAll	= `ERR_MONGO_SELECT_ALL`		// code 17
	ErrMongoInsert		= `ERR_MONGO_INSERT`			// code 18
	ErrMongoUpdate		= `ERR_MONGO_UPDATE`			// code 19
	ErrMongoDelete		= `ERR_MONGO_DELETE`			// code 20
)
