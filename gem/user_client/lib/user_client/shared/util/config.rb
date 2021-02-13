# respected for kaminari gem
# https://github.com/amatsuda/kaminari/blob/master/lib/kaminari/config.rb

require 'active_support/configurable'

module UserClient
  class << self
    def configure(&block)
      yield @config ||= UserClient::Configuration.new
    end

    def config
      @config
    end
  end

  class Configuration
    include ActiveSupport::Configurable
    config_accessor :db_yml
  end

  configure do |config|
    config.db_yml = ''
  end
end
