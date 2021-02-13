class User < ApplicationRecord
  has_many :user_notes
  has_many :notes, through: :user_notes

  # respond for exclude :registerable
  attr_accessor :password

  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable
  devise :trackable, :validatable,
         :omniauthable, omniauth_providers: [:google_oauth2]

  scope :wrote_notes, -> (user_id) { find_by(id: user_id).notes.order(:created_at) }

  class << self
    def from_omniauth(access_token)
      data_info = access_token.info
      user      = User.where(email: data_info['email']).first

      unless user
          user = User.create(
            email: data_info['email'],
            name:  data_info['name'],
            password: Devise.friendly_token[0, 20]
          )
      end

      user
    end
  end
end
