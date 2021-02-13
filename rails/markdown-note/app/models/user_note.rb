class UserNote < ApplicationRecord
  belongs_to :user
  belongs_to :note

  class << self
    def written_by_user?(note_id, user_id)
      user_note = find_by(note_id: note_id)
      return false if user_note.blank?
      user_note.user.id == user_id
    end
  end
end
