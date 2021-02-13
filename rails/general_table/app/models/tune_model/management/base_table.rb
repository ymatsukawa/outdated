module TuneModel
  module Management
    class BaseTable < PreparedModel
      include ActiveModel::Model

      attr_accessor :description

      validates :description,
        presence: true,
        length: { maximum: 255 }

      class << self
        def new_instance(create_params = {})
          obj = self.new
          obj.description = create_params[:description]
          obj
        end
      end

      def base_table_descriptions
        PreparedModel.run(@con, C.sql.select.base_table_descriptions)
      end

      # MUST: seek all insertion and optimize error handling.
      def create_base_tables
        begin
          ActiveRecord::Base.transaction do
            PreparedModel.run(@con, C.sql.insert.base_table_descriptions, [@description])
          end
        rescue => e
          false
        end
        true
      end

    end
  end
end
