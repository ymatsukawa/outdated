# added request spec @ custom_v0.0.1
require 'rails_helper'

RSpec.describe 'Users', type: :request do
  include_context "hook 'before and after' process"
  include_context 'prepare stub sign_in/sign_out methods'

  When { sign_in FactoryGirl.create(:user) }

  describe 'PUT /users' do
    Given(:put_params) { VALID_USERS_REQUEST_PARAMS }
    When(:res) do
      put '/users', put_params
      response
    end

    context 'when request parameter is valid' do
      it_behaves_like 'return response status code with', 302
      Then { expect(location_of(res, :path)).to eq '/' }
    end

    context 'when request parameter is invalid' do
      INVALID_USERS_REQUEST_PARAMS_LIST.each do |resource, value|
        context "when resource is #{resource} and value is #{value}" do
          Given(:put_params) do
            alter_params(
              :user,
              VALID_USERS_REQUEST_PARAMS,
              { resource => value }
            )
          end
          it_behaves_like 'return response status code with', 200
        end
      end
    end
  end
end
