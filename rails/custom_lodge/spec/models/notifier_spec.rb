# added @ custom_v0.0.1
require 'rails_helper'
RSpec.describe Notifier, type: :model do
  describe '.send_notice_email' do
    When(:user) { FactoryGirl.create(:user) }
    When(:article) do
      comment = FactoryGirl.create(:comment)
      article = FactoryGirl.create(:article)
      article.comments << comment
      article
    end
    When(:notice_email_receivers) { User.notice_email_receivers }

    context 'when exist notice email receiver' do
      [
        ['created', 'article'],
        ['posted comment', 'comment']
      ].each do |create_or_comment_to, sub_resource|
        context "and #{create_or_comment_to} article" do
          When { Notifier.send_notice_email(article, sub_resource) }
          Then { expect(ActiveJob::Base.queue_adapter.enqueued_jobs.size).to eq notice_email_receivers.size }
        end
      end

      context 'and correct parameter passed' do
        When do
          User.where(notice_mail_send_state: 1).update_all(notice_mail_send_state: 0)
          FactoryGirl.create(:user)
        end
        Then do
          bodies = { article_id: article.id, subject: article.title, body: article.body }
          expect(Notify).to receive(:notice_email).with(notice_email_receivers.last.email, 'article', bodies).and_return(double(Notify, deliver_later!:  true))
          Notifier.send_notice_email(article)
        end
      end
    end

    context 'when not exist notice email receiver' do
      When do
        User.where(notice_mail_send_state: 1).update_all(notice_mail_send_state: 0)
      end
      When { Notifier.send_notice_email(article) }
      Then { expect(ActiveJob::Base.queue_adapter.enqueued_jobs.size).to eq 0 }
    end

    # see this is out of test case ... ymatsukawa's taste
    context 'when passed incorrect sub resource' do
      Given(:expected_error_message) { Settings.lodge.notifier.error.implementation_error }
      Given(:raise_implementation_error) { raise_error(LodgeErrors::ImplementationError, expected_error_message) }
      Then do
        expect { Notifier.send_notice_email(article, 'mistook subresource') }.to raise_implementation_error
      end
    end
  end
end
