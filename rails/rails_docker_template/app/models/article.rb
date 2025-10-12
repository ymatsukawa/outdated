class Article < ApplicationRecord
  scope :only_visible, -> { where(archived: false) }

  def self.create_record(attrs)
    title = attrs[:title]
    text = attrs[:text]
    begin
      create!({ text: text, title: title })
      return true
    rescue => e
      return false
    end
  end


  def self.archive(id)
    begin
      find_by(id: id).update(archived: true)
      return true
    rescue => e
      return false
    end
  end
end
