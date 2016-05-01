# added @ custom_v0.0.1
require 'rails_helper'

RSpec.describe Notify, type: :mailer do
  When(:article) { FactoryGirl.create(:article) }
  When { ActionMailer::Base.deliveries.clear }

  describe '.notice_email' do
    Given(:article_body) do
      Settings.test.notify.notice_email.article.body
    end
    Given(:comment_body) do
      Settings.test.notify.notice_email.comment.body
    end

    context 'when parameter passed' do
      [
        ['article', BODY],
        ['comment', COMMENT_BODY]
      ].each do |sub_resource, body|
        context "when sub resource is #{sub_resource}" do
          When(:email) do
            bodies = {
              article_id: article.id,
              subject:    article.title,
              body:       body
            }
            Notify.notice_email(EMAIL, sub_resource, bodies).deliver_now
          end

          Then { expect(ActionMailer::Base.deliveries.count).to eq 1 }
          Then do
            subject_head     = Settings.lodge.notify.notice_email.subject.head
            reply_title_head = Settings.lodge.comment.reply_title.head
            if sub_resource.eql?('article')
              expect(email.subject).to eq (subject_head + article.title)
            else
              expect(email.subject).to eq (subject_head + reply_title_head + article.title)
            end
          end
          Then { expect(email.body.raw_source).to include (ROOT_URI + '/articles/' + article.id.to_s) }
          Then { expect(email.body.raw_source).to include body }
          Then do
            if sub_resource.eql?('article')
              expect(email.body.raw_source).to include 'New article:'
            else
              expect(email.body.raw_source).to include 'Commented at:'
            end
          end
          Then { expect(email.from.first).to eq Settings.test.notify.from }
          Then { expect(email.to.first).to eq EMAIL }
        end
      end
    end
  end
end
