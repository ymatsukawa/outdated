import SlackMessage                          from '../models/slack_message'
import HttpClient                            from '../models/http_client'
import { SCHEME, DOMAIN, PATH }              from '../const/request'
import { STATUS_CODE, successRes, errorRes } from '../const/response'

export const createTicket = (message) => {
  const slackMessage = new SlackMessage(message)
  if(!slackMessage.validate()) {
    return errorRes()
  }
  return requestCreateTicket(slackMessage)
}

const requestCreateTicket = (slackMessage) => {
  const httpClient = new HttpClient(slackMessage.getRequestUrl(), slackMessage.getMessageForRequest())
  const response   = httpClient.post(PATH.CREATE_TICKET)
  if(response.statusCode === STATUS_CODE.SUCCESS) {
    return successRes(response.message)
  } else {
    return errorRes()
  }
}
