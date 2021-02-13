module UserClient
  module Models
    SETTING = UserClient::Shared::Util::Setting

    class User < ActiveRecord::Base
      validates :email, :name, :password, presence: true
      validates_format_of :name, :password, with: /#{SETTING.data['model']['format']['not_identifier']}/
      validates_length_of :email, :name, maximum: 255
      validates_length_of :password, within: 8..255

      def update_record!(type_and_value)
        update_attributes!(type_and_value)
        self.touch
        self
      end
    end
  end
end
