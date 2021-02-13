# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'user_client/version'

Gem::Specification.new do |spec|
  spec.name          = 'user_client'
  spec.version       = UserClient::VERSION
  spec.authors       = ['Yusuke Matsukawa']
  spec.email         = ['mat27y@gmail.com']
  spec.summary       = %q{communiate to Rails activerecord like HTTP}
  spec.description   = %q{communiate to Rails activerecord like HTTP}
  spec.homepage      = ''
  spec.license       = 'MIT'

  spec.files         = `git ls-files -z`.split("\x0")
  spec.executables   = spec.files.grep(%r{^bin/}) { |f| File.basename(f) }
  spec.test_files    = spec.files.grep(%r{^(test|spec|features)/})
  spec.require_paths = ['lib']

  spec.add_dependency 'sqlite3'
  spec.add_dependency 'railties'
  spec.add_development_dependency 'activesupport'
  spec.add_development_dependency 'activerecord'
  spec.add_development_dependency 'bundler', '~> 1.6'
  spec.add_development_dependency 'factory_girl_rails'
  spec.add_development_dependency 'rake'
  spec.add_development_dependency 'rspec'
  spec.add_development_dependency 'simplecov'
  spec.add_development_dependency 'pry-byebug'
end
