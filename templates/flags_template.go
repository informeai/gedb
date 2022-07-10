package templates

var (
	Header = `
-----GEDB------
Describe:
gedb is utility for export data from 
databases to files in formats.(eg: json,csv)

`

	Usage = `
usage: gedb [COMMAND] [DRIVER] [-host] [-u <username>] [-p <password>]
[-db <database>] [-P <port>] [-f <format>] [-q <query>]

These are drives:
- mysql

These are common gedb commands used:
export		export tables for files in actual directory.
`
)
