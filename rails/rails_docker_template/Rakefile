require_relative "config/application"

Rails.application.load_tasks

namespace :ridge do
  desc 'reset db'
  task :reset do
    tasks = ['db:drop', 'db:create', 'ridge:apply']
    run_task(tasks)
  end

  desc 'seed db'
  task :seed do
    tasks = ['ridge:reset', 'db:seed']
    run_task(tasks)
  end

  task :apply do
    ENV['RAILS_ENV'] ||= 'development'
    exec_ridgepole('--apply', "-E #{ENV['RAILS_ENV']}")
  end
end

def run_task(tasks)
  tasks.each { |t| Rake::Task[t].invoke }
end

def exec_ridgepole(*options)
  command = ['bundle exec ridgepole', '-c config/database.yml', '-f db/Schemafile']
  system [command + options].join(' ')
end