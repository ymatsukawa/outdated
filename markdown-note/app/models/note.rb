class Note < ApplicationRecord
  has_many :user_notes, dependent: :destroy
  has_many :note_tags,  dependent: :destroy

  has_many :users, through: :user_notes
  has_many :tags,  through: :note_tags
end
