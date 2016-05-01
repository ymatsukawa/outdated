# Read about factories at https://github.com/thoughtbot/factory_girl

# mod @ custom_v0.0.1
FactoryGirl.define do
  factory :comment do
    body "dummy body"
    user
    article
  end
end
