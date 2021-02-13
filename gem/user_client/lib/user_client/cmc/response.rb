require 'active_support'
require 'user_client/shared/util/setting'

module UserClient
  module CMC
    class Response
      SETTING = UserClient::Shared::Util::Setting
      class << self
        def format_hashnize(record)
          return true if record == true # respond to #delete request
          return nil if record.blank?
          response = { 'body' => {} }
          SETTING.data['response']['respond_column'].each do |column|
            response['body'].merge!(column => record[column])
          end
          response.with_indifferent_access
        end
      end
    end
  end
end
