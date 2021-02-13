module UserClient
  module Validation
    module Models
      extend self
      def raise_if_invalid_entity(error)
        case error
        when ActiveRecord::RecordInvalid, ActiveRecord::RecordNotFound
          raise UserClient::BadRequest::ContentError
        else
          nil
        end
      end
    end
  end
end
