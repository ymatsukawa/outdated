require 'slack-ruby-client'
require_relative './http_client/request'
require_relative './http_client/request_service/create_ticket'
require_relative './http_client/request_service/project_info'
require_relative './http_client/request_service/error'

Slack.configure do |config|
  config.token = ENV['SLACK_BOT_API_TOKEN'] 
end

client = Slack::RealTime::Client.new

client.on :message do |message|
  request_service =
    case(message['text'].split("\n")[1]) # get slack text at second line
    when 'give info'
      HttpClient::RequestService::ProjectInfo.new({ message: message['text'] })
    when 'create ticket'
      HttpClient::RequestService::CreateTicket.new({ message: message['text'] })
    else
      HttpClient::RequestService::Error.new
    end
  response_message = HttpClient::Request.new(request_service).request[:message]
  client.message(channel: message['channel'], text: response_message)
end

client.start!
