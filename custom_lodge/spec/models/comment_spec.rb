require 'rails_helper'

RSpec.describe Comment, type: :model do
  it { should belong_to :user }
  it { should belong_to :article }
  it { should validate_presence_of(:body) }

  # mod @ custom_v0.0.1
  Given(:user)    { FactoryGirl.create(:user)    }
  Given(:article) { FactoryGirl.create(:article) }
  Given(:comment) { FactoryGirl.create(:comment) }
  Given(:expected_title) do
    Settings.lodge.comment.reply_title.head + comment.article.title
  end

  # added @ custom_v0.0.1
  describe '#reply_title' do
    When(:reply_title) { comment.reply_title }
    Then { expect(comment.reply_title).to eq expected_title }
  end
end
