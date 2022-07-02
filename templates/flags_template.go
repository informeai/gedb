package templates

var (
	Header = `
-----GEDB------
Describe:
gedb is utility for export data from 
databases to files in formats.(eg: json,csv)

`

	Usage = `
usage: gedb [DRIVER] [COMMAND] [-host] [-user <username>] [-pass <password>]
[-db <database>]

These are drives:
- mysql

These are common gedb commands used:
export		export tables for files in actual directory.
`
)
