# added @ custom_v0.0.1
require 'rails_helper'

module SpecDao
  RSpec.describe Dao::User, type: :model do
    When(:user) { FactoryGirl.create(:user) }
    def receivers(state)
      # must specified :id field because activejob needed 'identified' field
      Dao::User.select(:id, :email).where(notice_mail_send_state: state)
    end

    context '.notice_mail_send_state_where' do
      [0, 1].each do |state|
        context "when notice_mail_send_state is #{state}" do
          When(:expected_records) { receivers(state) }
          When(:returned_records) do
            Dao::User.notice_mail_send_state_where(state)
          end
          Then { expect(returned_records).to eq expected_records }
        end
      end
    end
  end
end
