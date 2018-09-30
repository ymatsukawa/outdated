module TuneModel
  module Management
    class CustomField < PreparedModel
      include ActiveModel::Model

      attr_accessor :base_table_id
      attr_accessor :description

      validates :base_table_id,
        presence: true,
        numericality: { only_integer: true, greater_than_or_equal_to: 1 }

      validates :description,
        presence: true,
        length: { maximum: 255 },
        if: -> { not self.description.nil? }

      class << self
        def new_instance(params = {})
          obj = self.new
          obj.base_table_id = params[:base_table_id]
          obj.description   = params[:description]
          obj
        end
      end

      def custom_field_descriptions
        return [] if invalid?
        PreparedModel.run(@con, C.sql.select.custom_field_descriptions, [@base_table_id]).to_a
      end

      def create_custom_field
        return false if invalid?
        begin
          ActiveRecord::Base.transaction do
            result_set = PreparedModel.run(@con, C.sql.insert.custom_field_descriptions, [@description]).first
            PreparedModel.run(@con, C.sql.insert.base_table_custom_fields, [@base_table_id, result_set['custom_field_id']])
          end
        rescue => e
          false
        end
        true
      end

    end
  end
end
