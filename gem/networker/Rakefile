# frozen_string_literal: true

require "bundler/gem_tasks"
require "rspec/core/rake_task"

RSpec::Core::RakeTask.new(:spec)

require "rubocop/rake_task"

RuboCop::RakeTask.new

task default: %i[spec rubocop]

namespace :test do
  desc "run mock server"
  task :server do
    ruby "spec/test_server/main.rb -p 3000 -o 0.0.0.0"
  end
end