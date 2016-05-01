require 'rails'
require 'user_client/model'

module UserClient
  class Railtie < Rails::Railtie
    initializer 'initialize database connection' do
      UserClient.config.db_yml = Rails.root.to_s + '/config/database.yml'
      Model.init
    end
  end
end
