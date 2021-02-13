# UserClient

Communiate DB(table:users) in Rails like HTTP request

## Given

You have

* Ruby on Rails 4
* db/development.sqlite3
* DB has table 'users', fields have [id, email, name, password, updated_at, created_at]

## Description

At **{YOUR RAILS_ROOT APP}/models/client.rb**

```
class Client
  include UserClient::Base
  
  def i_want_to_get_user_with_id(id)
    get_res = get(id) #=> RESPONSE
  end
  
  def i_want_to_get_user_with_email(email)
    get_res = get(email) #=> RESPONSE
  end

  def i_want_to_update_user_with(identifier, properties_and_entities)
    put(identifier, properties_and_entities) #=> RESPONSE with updated entity
  end

  def i_want_to_create_user_with(properties_and_entities)
    post(properties_and_entities) #=> RESPONSE with created entities
  end

  def i_want_to_delete_user_with(id)
    delete(id) #=> RESPONSE with deleted entities
  end
end
```

#### RESPONSE
```
res #=>
{
  "body" => {
              "id"=>1,
              "email"=>"dummy@dummy.com",
              "name"=>"dummy",
              "password"=>"dummy_password",
              "created_at"=>2014-11-08 14:20:52 UTC,
              "updated_at"=>2014-11-08 14:20:52 UTC
             }
}

# response can be approached with indifferent access
res[:body][:id]   #=> 1
res['body']['id'] #=> 1
```

## Detail

#### UserClient::Base#get

```
UserClient::Base#get(identifier)

UserClient::Base#get(1)
UserClient::Base#get('1')
UserClient::Base#get('dummy@dummy.com')
```

* identifier receives id or email

#### UserClient::Base#post


```
UserClient::Base#post(property_entity)

UserClient::Base#post({ email: 'foo@dummy.com', name: 'my_name', password: 'my_password' })
```

* email, name and password you want to create should be set as hash on second arg

#### UserClient::Base#put

```
UserClient::Base#put(identifier, property_entity)

UserClient::Base#put(1,                 { email: 'my_email_address@dummy.com' })
UserClient::Base#put('dummy@dummy.com', { email: 'foo@dummy.com', name: 'my_name' })
```

* identifier receives id or email
* email, name or password you want to update should be set as hash on second arg

#### UserClient::Base#delete

```
UserClient::Base#delete(identifier)

UserClient::Base#delete(1)
UserClient::Base#delete('1')
UserClient::Base#delete('dummy@dummy.com')
```

* identifier receives id or email

## Constraints

##### Rails app should

* set RDBMS as 'sqlite3'

##### RDBMS should

* have 'users' table
* 'users' table design should
  * have property
    * id, email, name, password, updated_at, created_at

## Installation

Given: Not pushed rubygems

#### with {YOUR RAILS APP ROOT}/Gemfile

```
$ git clone {user_client repository}
$ mv user_client {YOUR RAILS APP ROOT}/lib

-- {YOUR RAILS APP ROOT}/Gemfile
gem 'user_client', path: 'lib/user_client'

$ bundle install
```
