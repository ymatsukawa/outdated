class CustomDevise::RegistrationsController < Devise::RegistrationsController
  before_action :customize_sign_up_params, only: [:create, :update]

  def destroy
    @user.lock_access!
    redirect_to root_path
  end

  private

  def customize_sign_up_params
    [:name, :notice_mail_send_state].each do |add_field| # added flag field @ custom_v0.0.1
      devise_parameter_sanitizer.for(:account_update) << add_field
    end
    devise_parameter_sanitizer.for(:sign_up) << :name
  end

  def after_inactive_sign_up_path_for(resource)
    new_user_session_path
  end

end
