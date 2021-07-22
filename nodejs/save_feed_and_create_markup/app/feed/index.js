// https://node-postgres.com/features/queries
var FeedParser = require('feedparser')
var fetch = require('node-fetch')
const { getFeedCategory } = require('../common/feed_category')
const DBClient = require('../db/client')
const SQL = require('../db/sql')

/* WANT
create as class and promise base
*/
const createFeed = (category_type, feedLink) => {
  const feedparser = new FeedParser()
  const dbClient = new DBClient()
  console.log(feedLink)
  const req = fetch(feedLink)

  console.log("fetching target feed")
  req.then((res) => {
    if (res.status == 200) {
      console.log("got target feed")
      console.log(feedparser)
      res.body.pipe(feedparser)
      console.log("piped body to feedparser")
    }
    else {
      console.log('ERROR: response has bad status code - ' + res.status)
      dbClient.disconnect()
    }
  }, (error) => {
    console.log('ERROR: handled error of feed request')
    console.log(error)
    dbClient.disconnect()
  })

  feedparser.on('error', (error) => {
    console.log('ERROR: on feedparser')
    console.log(error)
  })

  // WANT: before insert, get all of category
  // now, all illegal data is violated by db.
  feedparser.on('readable', () => {
    var stream = feedparser
    var item = null
    
    var categoryNum = getFeedCategory(category_type)
    var title = ""
    var link = ""
    console.log("will parse and insert items")
    while (item = stream.read()) {
      title = sanitizeTitle(item.title)
      link = removeQueries(item.link)
      const sql = [SQL.feed.insert.got_feed, " ('", title, "', '", link, "', ", categoryNum.toString(), ", NOW())"].join("")
      console.log(sql)
      dbClient.insert(sql)
        .then((isInserted) => {
          // do what you want
        })
        .catch((error) => {
          // error is traced in db client
        })
    }

  })

  feedparser.on('end', () => {
    dbClient.disconnect()
  })
}

const sanitizeTitle = (title) => {
  return title.replace(/&/g, "&amp;")
              .replace(/"/g, "&quot;")
              .replace(/'/g, "&apos;")
              .replace(/</g, "&lt;")
              .replace(/>/g, "&gt;")
}

const removeQueries = (hyperlink) => {
  return hyperlink.replace(/\?.*$/, "")
}

exports.createFeed = createFeed