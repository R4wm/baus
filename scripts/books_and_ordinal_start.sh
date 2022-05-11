#!/bin/bash

for i in $(sqlite3 ~/kjv.db "select distinct(book) from kjv")
do
    ordinalStart=$(sqlite3 ~/kjv.db "select ordinal_verse from kjv where book=\"$i\" and chapter=1 and verse=1")
    echo "\"$i\": $ordinalStart,"
done

