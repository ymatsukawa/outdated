SOURCE_DB=db.sqlite.example
TARGET_DB=db.sqlite

if [ -e $TARGET_DB ]; then
  exit
fi

cp $SOURCE_DB $TARGET_DB

sqlite3 $TARGET_DB <<EOF
CREATE TABLE feeds (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  link TEXT NOT NULL UNIQUE,
  category TEXT NOT NULL
)
EOF

