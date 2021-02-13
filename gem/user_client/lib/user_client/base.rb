require 'user_client/validation/requests'
require 'user_client/validation/models'
require 'user_client/cmc/request'
require 'user_client/cmc/response'

module UserClient
  module Base
    include UserClient::BadRequest

    def get(id)
      get_validate!(id)
      response_on_request(:get_record, format_identifier(id))
    end

    def post(properties_and_entities)
      post_validate!(properties_and_entities)
      response_on_request(:create_record, properties_and_entities)
    end

    def put(id, properties_and_entities)
      put_validate!(id, properties_and_entities)
      response_on_request(:update_record, format_identifier(id).merge(properties_and_entities))
    end

    def delete(id)
      delete_validate!(id)
      response_on_request(:destroy_record, format_identifier(id))
    end

    private

    def format_identifier(id)
      id_regexp = /\A#{SETTING.data['request']['format']['id']}\z/
      # convert object class from Fixnum to String to respond Fixnum id
      id.to_s =~ id_regexp ? { id: id } : { email: id }
    end

    def response_on_request(request, properties_and_entities)
      begin
        response = UserClient::CMC::Request.send(request, properties_and_entities)
        UserClient::CMC::Response.format_hashnize(response)
      rescue => e
        UserClient::Validation::Models.raise_if_invalid_entity(e)
        raise e.class
      end
    end
  end
end
