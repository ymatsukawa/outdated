# frozen_string_literal: true

module Tablesyntax
  module Formatter
    class TableFormat
      def self.as_table(maps)
        keys = maps["keys"]
        table_kvs = maps["table_kvs"]

        keys_with_comma = keys.map { |s| ":#{s}" }.join(", ")
        body = create_tables(table_kvs)

        ["where(#{keys_with_comma}) do", body, "end\n"].join("\n")
      end

      def self.create_tables(table_kvs)
        table_vals = table_kvs.map(&:values).flatten
        col_max_width = table_vals.map(&:to_s).max_by(&:length).length

        table_kvs.map { |kvs| kvs.values.map { |v| format_obj(v, col_max_width) }.join(" | ") }
      end

      def self.format_obj(obj, width)
        case obj
        when String
          if exclude_as_string?(obj)
            obj.to_s
          else
            format('"%s"', obj)
          end
        when Symbol
          format(":%s", obj)
        when NilClass
          "nil"
        else
          obj.to_s
        end.rjust(width)
      end

      def self.exclude_as_string?(str)
        str.nil? || str == "nil" || /\s+\*\s+/.match?(str)
      end

      private_class_method :create_tables
      private_class_method :format_obj
      private_class_method :exclude_as_string?
    end
  end
end
