require 'active_support/core_ext'
require 'user_client/shared/util/setting'

module UserClient
  module Validation
    module CommonValidation
      SETTING            = UserClient::Shared::Util::Setting
      PERMITTED_CLASS    = SETTING.data['request']['permitted_class'].map { |klass| klass.constantize }
      PERMITTED_PROPERTY = SETTING.data['request']['permitted_property'].map { |property| property.to_sym }

      def identifier_validate!(id)
        raise BadRequest::IdentifierError unless as_id?(id) or as_email?(id)
      end

      def type_validate!(obj)
        raise BadRequest::TypeError unless PERMITTED_CLASS.include?(obj.class)
      end

      def property_validate!(property)
        raise BadRequest::PropertyError unless PERMITTED_PROPERTY.include?(property)
      end

      def as_id?(id)
        return true if id.to_s =~ /\A#{SETTING.data['request']['format']['id']}\z/
        false
      end

      def as_email?(email)
        return true if email =~ /\A#{SETTING.data['request']['format']['email']}\z/
        false
      end
    end
  end
end
