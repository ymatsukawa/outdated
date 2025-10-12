# frozen_string_literal: true

module Tablesyntax
  module Parser
    class OptionParser
      def initialize(options)
        @options = options
        @valid = true
        @key = nil
        @val = nil
      end

      module Option
        FILE_READ = :file
      end

      def parse
        if options.length.zero?
          @valid = false
          return
        end

        key, val = options
        if key_or_val_invalid?(key, val)
          @valid = false
          return
        end

        @key = key
        @val = val
      end

      def valid?
        @valid
      end

      attr_reader :key, :val

      private

      VAL_REQUIRED = :required
      ACCEPTED_OPTIONS = {
        Option::FILE_READ => VAL_REQUIRED
      }.freeze

      def key_or_val_invalid?(key, val)
        !ACCEPTED_OPTIONS.include?(key) || (ACCEPTED_OPTIONS[key] == VAL_REQUIRED && val.nil?)
      end

      attr_reader :options, :valid
    end
  end
end
