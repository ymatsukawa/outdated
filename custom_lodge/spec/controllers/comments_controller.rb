# added @ custom_v0.0.1
require 'rails_helper'

RSpec.describe CommentsController, type: :controller do
  Given(:user)    { FactoryGirl.create(:user) }
  Given(:article) { FactoryGirl.create(:article) }
  When { sign_in user }

  describe '#create' do
    [
      ['valid',   {}],
      ['invalid', { body: BLANK }]
    ].each do |valid_or_invalid, alter_entity|
      context "when passed #{valid_or_invalid} params" do
        When(:post_params) do
          alter_params(
            :comment,
            VALID_COMMENTS_CONTROLLER_PARAMS,
            { user_id: user.id, article_id: article.id }.merge(alter_entity)
          )
        end
        When { post :create, post_params }
        Then do
          if valid_or_invalid.eql?('valid')
            expect(flash.now['notice']).to eq 'Comment was successfully created.'
          else
            expect(flash.now['alert']).to  eq 'Failed to create comment.'
          end
        end
      end
    end
  end
end
