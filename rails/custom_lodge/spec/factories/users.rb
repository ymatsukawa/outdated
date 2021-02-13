# Read about factories at https://github.com/thoughtbot/factory_girl

FactoryGirl.define do
  sequence(:name) {|n| "John Doe #{n}" }
  sequence(:email) {|n| "john.doe.#{n}@example.com" }

  factory :user do
    name
    email
    password "password"
    notice_mail_send_state 1 # added @ custom_v0.0.1
    confirmed_at Time.zone.now
  end
end
