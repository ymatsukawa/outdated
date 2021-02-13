module UserClient
  module CMC
    class Request
      class << self
        def get_record(property_and_entity)
          UserClient::Model::User.find_by!(property_and_entity)
        end

        def create_record(properties_and_entities)
          record = UserClient::Model::User.new(properties_and_entities)
          return record if record.save!
        end

        def update_record(id_property_entity)
          record = get_record(id_property_entity[:id])
          if record.present?
            properties_and_entities = id_property_entity.reject { |key,value| key == :id }
            return record if record.update_record!(properties_and_entities)
          end
        end

        def destroy_record(property_and_entity)
          record = get_record(property_and_entity)
          return true if record.destroy!
        end
      end
    end
  end
end
