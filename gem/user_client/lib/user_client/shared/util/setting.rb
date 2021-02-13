require 'yaml'

module UserClient
  module Shared
    module Util
      class Setting
        attr_reader :data

        class << self
          def data
            @data ||= UserClient::Shared::Util::Setting.new.data
          end
        end

        def initialize
          @data = YAML.load_file(File.join(File.dirname(__FILE__), '../resource/setting.yml'))
        end
      end
    end
  end
end
