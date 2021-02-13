module Authentication
  extend ActiveSupport::Concern

  included do
    before_action :check_user_authentication

    def check_user_authentication
      if current_user.nil?
        redirect_to login_path
      end
    end
  end
end
