require 'rails_helper'
require 'json'

describe ApplicationController do
  # stub at external api request
  let(:api_request_header) { Stub::ExternalApi::Request.header(content_type) }
  let(:api_request_body) do
    { dummy_query: 'dummy' }
  end
  describe "request /dummy" do
    let(:path) { '/dummy' }

    where(:request_method) do
      [
        [:get], [:post]
      ]
    end
    with_them do
      where(:content_type) do
        [
          [:form_encode], [:json]
        ]
      end
      with_them do
        context 'when request to path where does not exist' do
          before do
            body = (content_type == :json ? api_request_body.to_json : api_request_body)
            send(request_method, '/dummy', params: body, headers: api_request_header)
          end
          it { expect(response.body).to eq Stub::ExternalApi::Response::NOT_FOUND_ERROR.to_json }
        end

      end
    end

  end
end
