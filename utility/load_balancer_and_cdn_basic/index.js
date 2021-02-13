const express = require('express')
const path = require('path')
const app = express()
const port = 3000

app.get('/', (req, res) => {
  console.log('accessed')
  res.send('hello')
})

app.get('/watch_video', (req, res) => {
  res.sendFile(path.join(__dirname, 'views/video_swan.html'))
})

app.listen(port, () => {
  console.log(`APP listening at http://localhost:${port}.`)
  console.log("HTTP REQUEST is listeing on http://localhost:8000")
})

