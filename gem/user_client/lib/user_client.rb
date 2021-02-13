if not defined?(Rails)
  $stderr.puts <<EOL
Rails frame work does not detected.

You should launch gem like below.

---
on RAILS_ROOT

$ rails c
> require 'user_client'

or

On RAILS_ROOT/models/dummy.rb
 class Dummy
   require 'user_client'
 end
EOL
  exit
end

require 'user_client/base'
require 'user_client/railtie'
require 'user_client/version'
