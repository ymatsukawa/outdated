module TuneModel
  class Table < PreparedModel
    include ActiveModel::Model

    attr_accessor :base_table_id
    attr_accessor :custom_field_ids

    validates :base_table_id,
      numericality: { only_integer: true, greater_than_or_equal_to: 0 },
      if: -> { not self.base_table_id.blank? }

    class << self
      def new_instance(params)
        obj = self.new
        obj.base_table_id    = params[:base_table_id]
        obj.custom_field_ids = []
        obj
      end
    end

    def specific_datas(base_table_id)
      return [] if invalid?
      base_table_id = base_table_id.present? ? base_table_id : -1
      PreparedModel.run(@con, C.sql.select.specific_table_datas, [base_table_id]).to_a
    end

    def select_box_base_table_ids
      result = PreparedModel.run(@con, C.sql.select.base_table_descriptions).to_a
      return [] if result.to_a.blank?
      result.map { |res| [res['description'], res['base_table_id']] }
    end

    def oldest_base_table_id
      return [] if invalid?
      result = PreparedModel.run(@con, C.sql.select.oldest_base_table_id).first
      return [] if result.blank?
      result['base_table_id']
    end

    def create_custom_field(params)
      exist_custom_field_ids = PreparedModel.run(@con, C.sql.select.custom_field_descriptions, [@base_table_id])
                                            .map { |record| record['custom_field_id'].to_s }
      return [] if (cids = get_custom_field_ids(params, exist_custom_field_ids)).blank?

      general_tables = []
      cids.each_with_index do |cid, index|
        general_tables << {
          base_table_id: @base_table_id, custom_field_id: cid,
          field_content: params["custom_field_id_#{cid}"]
        }
      end
      ::ActiveRecord::GeneralTable.import general_tables
    end

    private

    # MUST: With ActieModel Validation?
    def get_custom_field_ids(custom_field_ids_params, exist_custom_field_ids)
      ids = []
      custom_field_ids_params.each do |param_value|
        if param_value.match(/custom_field_id_[0-9]+/)
          cid = param_value.gsub(/custom_field_id_/, '')
          ids << cid if exist_custom_field_ids.include?(cid)
        end
      end
      ids
    end

  end
end
