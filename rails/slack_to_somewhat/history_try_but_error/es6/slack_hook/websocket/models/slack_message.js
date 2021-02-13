import { Record } from 'immutable'
import { SCHEME, DOMAIN, ENDPOINT } from '../const/request'

const attributes = {
  message:     "",
  projectId:   "",
  summary:     "",
  issueTypeId: "",
  priorityId:  ""
}

const SlackMessageRecord = Record(attributes)

export default class SlackMessage extends SlackMessageRecord {
  constructor(message) {
    super({message: message})
  }

  validate() {
    return true
  }

  getRequestUrl() {
    var domain = DOMAIN
    if(process.argv[2] === "production") {
      domain = (slackMessage.getProjectId() + domain)
    }
    return (SCHEME + domain + ENDPOINT)
  }

  getProjectId() {
    return "1"
  }


  getMessageForRequest() {
    return {
      projectId:   this.projectId,
      summary:     this.summary,
      issueTypeId: this.issueTypeId,
      priorityId:  this.priorityId
    }
  }
}
