require 'yaml'
require 'active_record'
require 'user_client/shared/util/config'
require 'user_client/models/user'

module UserClient
  module Model
    include UserClient::Models

    extend self
    def init
      db_config = YAML.load_file(UserClient.config.db_yml)
      ActiveRecord::Base.establish_connection(db_config[ENV['RAILS_ENV']])
      I18n.enforce_available_locales = true
    end
  end
end
