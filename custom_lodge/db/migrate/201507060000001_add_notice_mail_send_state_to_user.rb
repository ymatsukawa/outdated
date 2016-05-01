# added @ custom_v0.0.1
class AddNoticeMailSendStateToUser < ActiveRecord::Migration
  # RANGE_CHECK_AT_NOTICE_MAIL_SEND_STATE = 'notice_mail_send_state_range_check'
  def change
    # defined state
    #   0: not receive email any times
    #   1: receive email whenever article or comment post happened
    # not specified state
    #   2: receive email only when article post happened
    #   3: receive email only when comment post happened
    add_column :users, :notice_mail_send_state, :integer, null: false, default: 1

    # ---------------------------------------------------------------------------
    # Enable const RANGE_... and 'reversible' DSL if you use PostgreSQL
    # (SQLite3 and MySQL does not work because of RDBMS Specification)
    # Intent:
    #   App layer validation is to notice 'invalid input' error message for user.
    #   Data layer validation(CHECK constraint) is to protect data isolation.
    #
    # reversible do |version| 
    #   version.up do
    #     execute <<-SQL
    #       ALTER TABLE users
    #       ADD CONSTRAINT #{RANGE_CHECK_AT_NOTICE_MAIL_SEND_STATE}
    #       CHECK(notice_mail_send_state IN(0, 1))
    #     SQL
    #   end
    #
    #   version.down do
    #     # at deletion
    #     execute <<-SQL
    #       ALTER TABLE users
    #       DROP CONSTRAINT #{RANGE_CHECK_AT_NOTICE_MAIL_SEND_STATE}
    #     SQL
    #   end
    # end
  end
end
