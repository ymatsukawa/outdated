import { createTicket } from './actions'
import { STATUS_CODE }  from './const/response'

const Botkit = require('botkit')

const controller = Botkit.slackbot(
  { debug: (process.argv[2] === "debug") }
)

controller.spawn(
  { token: process.env.SLACK_BOT_API_TOKEN }
).startRTM()

controller.hears('create', ['direct_mention'], (bot, message) => {
  if(createTicket(message).statusCode === STATUS_CODE.SUCCESS) {
    bot.reply(message, 'success. see ' + response.message)
  } else {
    bot.reply(message, "something went wrong.\ncheck your wrote format.\nif evertything is good, contact administrator or create issue.")
  }
})
