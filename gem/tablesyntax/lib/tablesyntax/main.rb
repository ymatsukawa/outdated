# frozen_string_literal: true

require_relative "parser/option_parser"
require_relative "facade/file_reader"

module Tablesyntax
  class Main
    def self.run(*options)
      opts = Parser::OptionParser.new(options)
      opts.parse
      raise StandardError unless opts.valid?

      case opts.key
      when Parser::OptionParser::Option::FILE_READ
        Facade::FileReader.run(opts.val)
      end
    end
  end
end
