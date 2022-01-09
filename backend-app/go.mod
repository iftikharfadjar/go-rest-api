module backend

replace backend/models => /workspace/go-rest-api/backend-app/models

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/pascaldekloe/jwt v1.10.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/net v0.0.0-20211005001312-d4b1ae081e3b // indirect
	golang.org/x/sys v0.0.0-20211004093028-2c5d950f24ef // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect

)
