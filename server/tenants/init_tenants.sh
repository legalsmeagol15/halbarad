#!/bin/bash
set -e
sqlite3 -version

function do_cmd()
{
    echo "$1"
    sqlite3 tenants.db "$1" ;
}

echo "Configuring SQLite database..."

do_cmd "CREATE TABLE IF NOT EXISTS tenants (id INTEGER PRIMARY KEY, created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, username TEXT UNIQUE, pswdhash TEXT);"
do_cmd "INSERT OR IGNORE INTO tenants (username, pswdhash) VALUES ('test', 'testhash');"


echo "...SQLite database configuration complete."