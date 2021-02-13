module Management
  class CustomFieldsController < ApplicationController
    # MUST: prohibit access #index until base_table is created.
    def index
      @custom_fields = TuneModel::Management::CustomField.new_instance(index_params).custom_field_descriptions
    end

    def create
      unless TuneModel::Management::CustomField.new_instance(create_params).create_custom_field
        return redirect_to management_custom_fields_path
      end
      redirect_to management_custom_fields_path
    end

    # MUST create #edit, #update, #delete

    private

    def index_params
      {
        base_table_id: custom_field_params[:base_table_id]
      }
    end

    def create_params
      {
        base_table_id: custom_field_params[:base_table_id],
        description:   custom_field_params[:description] || ''
      }
    end

    def custom_field_params
      params.permit(C.permitted_attributes.management.custom_fields)
    end
  end
end
