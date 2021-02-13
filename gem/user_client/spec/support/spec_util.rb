def create_user
  UserClient::Model::User.new(
    id:       USER_ID,
    email:    EMAIL,
    name:     NAME,
    password: PASSWORD,
  ).save
end

def delete_user
  UserClient::Model::User.delete_all
end
