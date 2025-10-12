require 'faker'

records = []
10.times do
  records.push({
    title: Faker::Book.title,
    text: Faker::Lorem.sentence
  })
end
Article.create!(records)