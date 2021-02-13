module CustomDevise::RegistrationsHelper
  # added @ custom_v0.0.1
  def notice_mail_send?
    current_user.notice_mail_send_state.equal?(1)
  end
end
