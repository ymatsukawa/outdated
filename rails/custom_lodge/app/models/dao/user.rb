# added @ custom_v0.0.1
module Dao
  class User < ActiveRecord::Base
    class << self
      def notice_mail_send_state_where(state)
        self.select(:id, :email).where(notice_mail_send_state: state) 
      end
    end
  end
end
