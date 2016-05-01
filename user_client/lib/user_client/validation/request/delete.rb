require 'user_client/error'
require 'user_client/validation/common_validation'

module UserClient
  module BadRequest
    include UserClient::Validation::CommonValidation

    def delete_validate!(id)
      [:type_validate!, :identifier_validate!].each do |method|
        send(method, id)
      end
    end
  end
end
