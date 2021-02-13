ID_FORMAT           = '\d+'
EMAIL_FORMAT        = '[^@]+@([^@\.]+\.)+[^@\.]+'
PERMITTED_PROPERTY  = [:id, :email, :name, :password, :updated_at, :created_at]
IDENTIFIER_PROPERTY = [:id, :email]

class User
  include UserClient::Base
end
