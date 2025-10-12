require 'rails_helper'

RSpec.describe 'V1::CreateArticle', type: :request do
  include RequestSpecHelper

  describe 'POST /v1/article' do
    let(:params) do
      {
        article: {
          title: title,
          text: text
        }
      }
    end

    subject(:send_request) do
      post '/v1/articles', params: params
    end
    
    context 'when request is valid' do
      let(:title) { 'sample' }
      let(:text) { 'lorem ipsum' }

      it_behaves_like 'response with status', 201
    end

    context 'when request is invalid' do
      where(:title_, :text_) do
        [
          [ nil, 'a' ],
          [ 'a', nil ],
          [ [], 'a' ],
          [ 'a', [] ]
        ]
      end
      with_them do
        let(:title) { title_ }
        let(:text) { text_ }

        let(:expected_body) { V1::Common::Response.create_json(400) }
        it_behaves_like 'response with body and status', 400
      end
    end
  end
end