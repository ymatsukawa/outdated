FactoryBot.define do
  factory :article, class: Article do
    title { Faker::Book.title }
    text { Faker::Lorem.sentence }
    archived { false }
  end

  trait(:archive) do
    archived { true }
  end
end