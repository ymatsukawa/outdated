# frozen_string_literal: true

require "psych"
require "json"

require_relative "../fileio/output/standard_output"
require_relative "../formatter/table_format"

module Tablesyntax
  module Facade
    class FileReader
      def self.run(path)
        s = stream(path)
        Fileio::Output::StandardOutput.out(s)
      end

      def self.stream(path)
        ext = File.extname(path)
        case ext
        when ".yml", ".yaml"
          yaml_stream = File.read(path)
          yaml = Psych.safe_load(yaml_stream, permitted_classes: [Symbol])
          Formatter::TableFormat.as_table(yaml)
        when ".json"
          json = JSON.parse(File.read(path))
          Formatter::TableFormat.as_table(json)
        else
          raise StandardError
        end
      end

      private_class_method :stream
    end
  end
end
