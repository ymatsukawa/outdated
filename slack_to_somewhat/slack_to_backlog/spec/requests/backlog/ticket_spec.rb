require 'rails_helper'
require 'json'

describe Backlog::TicketsController do
  let(:endpoint)       { Stub::Backlog::Api::Common::END_POINT }
  let(:query_api_key)  { "?apiKey=#{Stub::Backlog::Api::Common::API_KEY}" }

  where(:content_type) do
    [
      [:form_encoded], [:json]
    ]
  end
  with_them do
    describe 'POST /backlog/ticket' do
      # stub at external api request
      let(:api_request_header) { Stub::ExternalApi::Request.header(content_type) }
      let(:api_request_body) do
        { ticket: { summary: 'test', project_id: '1234', priority_id: '1', issue_type_id: '1' } }
      end

      # stub at backlog_kit faraday http request
      let(:backlog_api_path) { '/issues' }
      let(:backlog_request_headers) { Stub::Backlog::Api::Common::REQUEST_HEADER }

      context 'when request is valid' do
        before do
          stub_request(:post, "#{endpoint}#{backlog_api_path}#{query_api_key}")
            .with(headers: backlog_request_headers, body: api_request_body[:ticket].dup.camelize_keys!)
            .to_return(
              status: 201,
              body: Stub::Backlog::Api::Response::CREATE_ISSUES,
              headers: { 'Content-Type' => 'application/json' }
            )
          body = (content_type == :json ? api_request_body.to_json : api_request_body)
          post '/backlog/ticket', params: body, headers: api_request_header
        end
        it { expect(response.body).to eq Stub::ExternalApi::Response::SUCCESS.to_json }
      end

      context 'when request is invalid' do
        where(:summary, :project_id, :priority_id, :issue_type_id) do
          [
             [nil   , nil   , nil, nil],
             ['test', nil   , nil, nil],
             [nil   , '1234', nil, nil],
             ['test', '1234', nil, nil],
             [nil   , nil   , '1', nil],
             ['test', nil   , '1', nil],
             [nil   , '1234', '1', nil],
             ['test', '1234', '1', nil],
             [nil   , nil   , nil, '1'],
             ['test', nil   , nil, '1'],
             [nil   , '1234', nil, '1'],
             ['test', '1234', nil, '1'],
             [nil   , nil   , '1', '1'],
             ['test', nil   , '1', '1'],
             [nil   , '1234', '1', '1'],
          ]
        end
        with_them do
          let(:api_request_body) do
            { ticket: { summary: summary, project_id: project_id, priority_id: priority_id, issue_type_id: issue_type_id } }
          end

          before do
            body = (content_type == :json ? api_request_body.to_json : api_request_body)
            post '/backlog/ticket', params: body, headers: api_request_header
          end

          it { expect(response.body).to eq Stub::ExternalApi::Response::ERROR.merge(message: 'check whether attributes include blank input.').to_json }
        end
      end
    end

  end
end
