#!/bin/sh

METHOD="POST"
HOST="localhost:3000"

URL="/login"
FORM="LoginID=viper Password=12345"

http -v --form $METHOD $HOST$URL $FORM

