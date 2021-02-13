require 'pry'
require 'pathname'
require 'yaml'
require 'active_record'

ENV['GEM_ENV'] = 'test'

# prepare test environment
db_config = YAML.load_file(File.expand_path('./spec/support/dummy_app/config/database.yml'))
ActiveRecord::Base.establish_connection(db_config[ENV['GEM_ENV']])
I18n.enforce_available_locales = true

# mock Rails module for test
module Rails
  extend self
  def root
    'spec/support/dummy_app'
  end

  def env
    ENV['GEM_ENV']
  end
end

require 'user_client'
Dir[Pathname.pwd.join('spec/support/**/*.rb')].each { |f| require f }

require 'simplecov'
SimpleCov.start

RSpec.configure do |config|
  config.expect_with :rspec do |expectations|
    expectations.include_chain_clauses_in_custom_matcher_descriptions = true
  end

  config.mock_with :rspec do |mocks|
    mocks.verify_partial_doubles = true
  end

  config.order = :random
end
