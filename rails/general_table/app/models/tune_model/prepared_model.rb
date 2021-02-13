module TuneModel
  class PreparedModel
    attr_accessor :con

    def initialize
      @con = ActiveRecord::Base.connection
    end

    class << self
      # MUST: re-think con is needed? i understand this is class method
      def run(con, sql, params = [])
        if params.blank?
          con.execute(sql)
        else
          con.raw_connection.exec_params(sql, params)
        end
      end
    end
  end
end
