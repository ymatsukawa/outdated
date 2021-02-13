FactoryGirl.define do
  factory :user, class: UserClient::Model::User do
    sequence(:id) { |n| n }
    email    EMAIL
    name     NAME
    password PASSWORD
    created_at DateTime.now
    updated_at DateTime.now
  end
end
