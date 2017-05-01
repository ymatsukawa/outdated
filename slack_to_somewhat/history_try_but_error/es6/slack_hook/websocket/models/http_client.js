import { Record }               from 'immutable'
import { successRes, errorRes } from '../const/response'

const attributes = {
  request:  null,
  reqUrl:   "",
  reqBody:  {},
  response: null
}

const HttpClientRecord = Record(attributes)

export default class HttpClient extends HttpClientRecord {
  constructor(reqUrl, reqBody) {
    super({
      request: require('request'),
      reqUrl:  reqUrl,
      reqBody: reqBody,
    })
  }

  post(path) {
    var res = null
    this.request.post(this.reqUrl + path, this.reqBody, function(error, response, body) {
      if(!error) {
        res = successRes('success url here')
      } else {
        res = errorRes()
      }
    })
    const sleep = require('sleep')
    sleep.sleep(5)
    return res
  }
}
