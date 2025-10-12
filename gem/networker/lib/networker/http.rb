# frozen_string_literal: true

require_relative "./parser/option_parser"
require_relative "./parser/map_file"
require_relative "./facade/http"

module Networker
  class Http
    def self.req(*options)
      options = Networker::Parser::OptionParser.new(options)
      options.parse

      raise StandardError unless options.valid?

      case options.key
      when Networker::Parser::OptionParser::Option::FILE_READ
        http_opts = Networker::Parser::MapFile.read(options.val)
        Networker::Facade::Http.req(http_opts)
      end
    end
  end
end
