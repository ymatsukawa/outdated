module Management
  class BaseTablesController < ApplicationController
    def index
      @tables = TuneModel::Management::BaseTable.new_instance.base_table_descriptions
    end

    def new
    end

    def create
      unless TuneModel::Management::BaseTable.new_instance(create_params).create_base_tables
        # do something
      end
      redirect_to management_base_tables_path
    end

    # MUST: create #edit, #update, #delete

    private

    def create_params
      {
        description: management_table_params['description']
      }
    end

    def management_table_params
      params.permit(C.permitted_attributes.management.base_tables)
    end
  end
end
