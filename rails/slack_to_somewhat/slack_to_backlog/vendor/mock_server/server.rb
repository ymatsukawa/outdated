require 'sinatra'
require 'pry'
require_relative '../../spec/support/stub/backlog/api/response/create_issues'

before { logger.level = 0 }

not_found do
  { "statusCode": 404, "message": "resource not found at mock_server" }.to_json
end

post "/api/v2/issues", provides: :json do
  content_type :json
  status 201
  response.headers['Location'] = 'https://ymatsukawa.backlog.jp/view/dummyPorject-1'
  Stub::Backlog::Api::Response::CREATE_ISSUES
end
