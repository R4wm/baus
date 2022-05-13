#!/bin/bash

for i in $(sqlite3 ~/kjv.db "select distinct(book) from kjv")
do
    maxChapter=$(sqlite3 ~/kjv.db "select max(chapter) from kjv where book=\"$i\"")
    echo "\"{$i\"}: {\"chapters\": $maxChapter, ordinal},"
done
