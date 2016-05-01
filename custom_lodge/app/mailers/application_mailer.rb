# added @ custom_v0.0.1
class ApplicationMailer < ActionMailer::Base
  default from: ENV['MAIL_SENDER']
  layout 'mailer'
end
