# frozen_string_literal: true

require "json"

module Networker
  module Parser
    class MapFile
      def self.read(path)
        raise StandardError unless File.exist?(path)

        file_to_map(path)
      end

      def self.file_to_map(path)
        case File.extname(path)
        when ".yaml", ".yml"
          # nope in v0.0.2
        when ".json"
          JSON.parse(File.read(path))
        else
          raise StandardError
        end
      end

      private_class_method :file_to_map
    end
  end
end
