const SQL = {
    feed: {
        select: {
            periodic: "SELECT title, hyperlink FROM feeds WHERE category = $0 AND created_at > NOW() - interval '6 days' ORDER BY title"
        },
        insert: {
            got_feed: "INSERT INTO feeds (title, hyperlink, category, created_at) VALUES"
        }
    }
}

module.exports = SQL