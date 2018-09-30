class TablesController < ApplicationController
  # MUST: commonize #index and #new
  def index
    table = TuneModel::Table.new_instance(index_params)
    @select_table_ids = table.select_box_base_table_ids

    # transition from index...
    #   :base_table_id is blank
    #   so get base_table_id from top of select_box
    @base_table_id = index_params[:base_table_id].blank? ? @select_table_ids.first[1] : index_params[:base_table_id]
    @table_datas   = table.specific_datas(@base_table_id)
    @custom_fields = TuneModel::Management::CustomField.new(index_params.merge(base_table_id: @base_table_id)).custom_field_descriptions
  end

  # MUST: constrain #new, unable to access if select_table_ids does not exist
  #   means accessing #new is prohibited untill base_table and custom_field have relation.
  def new
    table = TuneModel::Table.new_instance(new_params)
    @select_table_ids = table.select_box_base_table_ids

    # transition from index...
    #   :base_table_id is blank
    #   so get base_table_id from top of select_box
    @base_table_id = new_params[:base_table_id].blank? ? @select_table_ids.first[1] : new_params[:base_table_id]
    @table_datas   = table.specific_datas(@base_table_id)
    @custom_fields = TuneModel::Management::CustomField.new(new_params.merge(base_table_id: @base_table_id)).custom_field_descriptions
  end

  def create
    return render :error if table_params[:base_table_id].blank? # MUST
    unless TuneModel::Table.new_instance(create_params).create_custom_field(params)
      return render :error
    end

    redirect_to action: :new, base_table_id: table_params[:base_table_id]
  end

  private

  def create_instance
    @table = Table.new_instance(table_params)
  end

  # MUST: commonize index_params and new_params
  def index_params
    {
      base_table_id: table_params[:base_table_id].present? ? table_params[:base_table_id] : ''
    }
  end

  def new_params
    {
      base_table_id: table_params[:base_table_id].present? ? table_params[:base_table_id] : ''
    }
  end

  def create_params
    {
      base_table_id: table_params[:base_table_id]
    }
  end

  def table_params
    params.permit(C.permitted_attributes.tables)
  end
end
