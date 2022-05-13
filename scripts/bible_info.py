#!/usr/bin/env python3

import sqlite3

con = sqlite3.connect('kjv.db')
cur = con.cursor()

books = ["GENESIS","EXODUS","LEVITICUS","NUMBERS","DEUTERONOMY","JOSHUA","JUDGES","RUTH","1SAMUEL","2SAMUEL","1KINGS","2KINGS","1CHRONICLES","2CHRONICLES","EZRA","NEHEMIAH","ESTHER","JOB","PSALMS","PROVERBS","ECCLESIASTES","SONG OF SOLOMON","ISAIAH","JEREMIAH","LAMENTATIONS","EZEKIEL","DANIEL","HOSEA","JOEL","AMOS","OBADIAH","JONAH","MICAH","NAHUM","HABAKKUK","ZEPHANIAH","HAGGAI","ZECHARIAH","MALACHI","MATTHEW","MARK","LUKE","JOHN","ACTS","ROMANS","1CORINTHIANS","2CORINTHIANS","GALATIANS","EPHESIANS","PHILIPPIANS","COLOSSIANS","1THESSALONIANS","2THESSALONIANS","1TIMOTHY","2TIMOTHY","TITUS","PHILEMON","HEBREWS","JAMES","1PETER","2PETER","1JOHN","2JOHN","3JOHN","JUDE","REVELATION"]


bible_data = []
for book in books:
    sqlstmt = 'SELECT ordinal_verse, max(chapter) from kjv where book=\'{}\''.format(book)
    for row in cur.execute(sqlstmt):
        bible_data.append({"book_name": book,
                            "ordinal_verse_start": row[0],
                            "chapters": row[1]})

print(bible_data)
