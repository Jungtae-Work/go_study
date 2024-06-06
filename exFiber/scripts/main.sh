#!/bin/sh

METHOD="GET"
HOST="localhost:3000"

URL="/"
SESSION_KEY="169ab5e3-a048-49b5-9320-2685a719c344"

http -v $METHOD $HOST$URL "Session-Id: $SESSION_KEY"

