# added @ custom_v0.0.1
require 'uri'

# usage:
#   resouce_name = :user
#   parameters   = { user: { user_id: 1, name: 'example', email: 'example@example.com' } }
#   alter_params = { name: 'sample', email: 'sample@example.com' }
#   > alter_params(resource_name, paramters, alter_params)
#   > #=> { user: { user_id: 1, name: 'sample', email: 'sample@example.com' } }
def alter_params(resource_name, parameters, alter_params)
  {
    resource_name =>
      parameters[resource_name].merge(alter_params)
  }
end

# usage:
#   res    = (Rspec HTTP Response)
#   detail = :scheme, :host, :port, :path or etc
#
# in Rspec
#   run  > get '/foo'
#   at   > scheme = http
#          host   = localhost
#          port   = 3000
#   then > res    = resposne
#
#   > location_of(res)
#   > #=> 'http://localhost:3000/foo'
#   > location_of(res, :path)
#   > #=> '/foo'

def location_of(res, detail = nil)
  uri = URI.parse(res.location)
  return uri if detail.nil?
  uri.send(detail)
end

def show_path_of(resource_name)
  /\/#{resource_name}\/(?=[^0])(?=\d{1,})/
end
