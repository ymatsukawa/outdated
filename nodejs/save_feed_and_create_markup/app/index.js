const express = require('express')
const app = express()
const { createFeed } = require('./feed/index')
const { createMarkup } = require('./file_maker/index')
const { getFeedLinks } = require('./common/feed_category')

app.use(express.json())

PORT = 3000

app.get('/', (req, res) => {
  json = {
    'hello': 'world'
  }

  res.json(json)
})

/*
  request is expectd
  Content-Type application/json
  json should include
  {
    category: "tech-all"
  }

  category type is explaind in common/feed_category.js
*/
app.post('/feed', (req, res) => {
  const reqJson = req.body
  console.log(reqJson)
  if (!req.is('application/json') || !reqJson || (typeof reqJson.category !== "string")) {
    res.json({
      code: 400,
      message: "content type or request format is invalid."
    });
    return
  }

  const feedLinks = getFeedLinks(reqJson.category)
  feedLinks.forEach(feedLink => {
    createFeed(reqJson.category, feedLink)
  })
  res.json({
    code: 201,
    message: "feed is created."
  })
})

/*
  request is expectd
  Content-Type application/json
  json should include
  {
    category: "tech-all"
  }

  category type is explaind in common/feed_category.js
*/
app.post('/feed/markup', (req, res) => {
  const reqJson = req.body
  console.log(reqJson)
  if (!req.is('application/json') || !reqJson || (typeof reqJson.category !== "string")) {
    res.json({
      code: 400,
      message: "content type or request format is invalid."
    });
    return
  }

  createMarkup(reqJson.category)
  res.json({
    code: 201,
    message: "feed is created."
  })
})

app.listen(PORT, () => {
  console.log('listening on: ' + PORT)
})
