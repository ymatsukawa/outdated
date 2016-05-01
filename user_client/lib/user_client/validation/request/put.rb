require 'user_client/error'
require 'user_client/validation/common_validation'

module UserClient
  module BadRequest
    include UserClient::Validation::CommonValidation

    def put_validate!(id, property_and_entity)
      [:identifier_validate!, :type_validate!].each do |method|
        send(method, id)
      end
      property_and_entity.each do |property, entity|
        property_validate!(property)
        type_validate!(entity)
      end
    end
  end
end
