require 'user_client/error'
require 'user_client/validation/common_validation'

module UserClient
  module BadRequest
    include UserClient::Validation::CommonValidation

    def post_validate!(properties_and_entities)
      properties_and_entities.each do |property, entity|
        property_validate!(property)
        type_validate!(entity)
      end
    end
  end
end
