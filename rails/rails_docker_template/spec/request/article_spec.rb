require 'rails_helper'

RSpec.describe "Article", type: :request do
  include RequestSpecHelper

  describe 'GET /articles' do
    subject(:send_request) do
      get '/articles'
    end

    it_behaves_like 'response with status code', 200
  end

  describe 'GET /articles/:id' do
    subject(:send_request) do
      get "/articles/#{id}"
    end

    context 'when id is valid' do
      before { create(:article) }
      let(:id) { Article.first.id }

      it_behaves_like 'response with status code', 200
    end

    context 'when id is invalid' do
      where(:id_) do
        [
          0, -1, 7.2, 'a', 'flock'
        ]
      end
      with_them do
        let(:id) { id_ }
        
        it_behaves_like 'response with status code', 500
      end
    end
  end

  describe 'GET /article/new' do
    pending
  end

  describe 'POST /article' do
    pending
  end

  describe 'PUT /article/:id' do
    pending
  end

  describe 'DELETE /article/:id' do
    pending
  end
end