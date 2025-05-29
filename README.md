# mytop-go 

Display MySQL server performance info like 'mytop'.

## Installation

`go get github.com/go-bridget/mytop-go/cmd/mytop@master`

## The Basics

mytop-go will connect to a MySQL server and periodically run the `SHOW PROCESSLIST` command and attempt to summarize the information from them in a useful format.

The program lists as many threads as can fit on screen. The display looks like:

![mytop-go display example](https://i.imgur.com/t0s5Ejp.png "mytop-go")

Often times the query info is what you are really interested in, so it is good to run mytop-go in an terminal that is wider than the normal 80 columns if possible.

## Arguments

mytop-go handles only short command-line arguments.

- **-u** Username to use when logging in to the MySQL server. Default: ''root''
- **-p** Password to use when logging in to the MySQL server. Default: none
- **-h** Hostname of the MySQL server. The hostname may be followed by an option port number. Note that the port is specified separate from the host when using a config file. Default: ''localhost''
- **-P** Port. If you're running MySQL on a non-standard port, use this to specify the port number. Default: 3306
- **-s** Determines delay between display refreshes. Default: 5
- **-d** Database. Use if you'd like mytop-go to connect to a specific database by default. Default: ''mysql''
- **-i** Idle.Specify if you want idle (sleeping) threads to appear in the list

## Shortcut Keys

The following keys perform various actions while mytop-go is running.

- **f** - Filter by query
- **u** - Filter by user
- **t** - Filter by time
- **s** - Change sort column
- **k** - Kill a thread by PID
- **K** - Kill all user read queries
- **q** - Stop mytop-go

## Contributors

- Tit Petric @titpetric
- Evan do Carmo @carmo-evan
