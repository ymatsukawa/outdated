# added request spec @ custom_v0.0.1
require 'rails_helper'

RSpec.describe 'Comments', type: :request do
  include_context "hook 'before and after' process"
  include_context 'prepare stub sign_in/sign_out methods'

  Given(:user)    { FactoryGirl.create(:user) }
  Given(:article) { FactoryGirl.create(:article) }
  When { sign_in user }

  describe 'POST /comments' do
    Given(:expected_location) { ARTICLE_SHOW_PATH_FORMAT }
    When(:res) do
      post '/comments', post_params
      response
    end

    ['valid', 'invalid'].each do |valid_or_invalid|
      context "when #{valid_or_invalid} parameter is passed" do
        Given(:post_params) do
          alter_resource = { user_id: user.id, article_id: article.id }
          alter_resource.merge!(body: BLANK) if valid_or_invalid.eql?('invalid')
          alter_params(
            :comment,
            VALID_COMMENTS_REQUEST_PARAMS,
            alter_resource
          )
        end

        it_behaves_like 'return response status code with', 302
        Then do
          expect(location_of(res, :path)).to match expected_location
        end
      end
    end
  end
end
