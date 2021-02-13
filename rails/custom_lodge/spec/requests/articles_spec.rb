# added request spec @ custom_v0.0.1
require 'rails_helper'

RSpec.describe 'Article', type: :request do
  include_context "hook 'before and after' process"
  include_context 'prepare stub sign_in/sign_out methods'

  Given(:user) { FactoryGirl.create(:user) }
  When { sign_in user }

  describe 'POST /article' do
    Given(:post_params) do
      alter_params(:article, VALID_ARTICLE_REQUEST_PARAMS,
                   { user_id: user.id })
    end
    When(:res) do
      post '/articles', post_params
      response
    end

    context 'when request parameter is valid' do
      it_behaves_like 'return response status code with', 302
      Then do
        expect(location_of(res, :path)).to match ARTICLE_SHOW_PATH_FORMAT
      end
    end

    context 'when request parameter is invalid' do
      INVALID_ARTICLE_REQUEST_PARAMS_LIST.each do |resource, value|
        context "when resource is #{resource} and value is #{value}" do
          Given(:post_params) do
            alter_params(:article, VALID_ARTICLE_REQUEST_PARAMS,
                         { resource => value })
          end
          it_behaves_like 'return response status code with', 200
        end
      end
    end
  end
end
