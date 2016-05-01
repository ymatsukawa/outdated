# added @ custom_v0.0.1
class Notify < ApplicationMailer
  SUBJECT_HEAD = Settings.lodge.notify.notice_email.subject.head
  REPLY_HEAD   = Settings.lodge.comment.reply_title.head

  def notice_email(to, resource, bodies)
    # use at views/notify/notice_email.text.erb
    @scheme   = Settings.lodge.global.uri.scheme
    @host     = Settings.lodge.global.uri.host
    @port     = Settings.lodge.global.uri.port
    @resource = resource
    @bodies   = bodies
    mail(
      to: to,
      subject: send_subject(resource, bodies[:subject])
    )
  end

  private

  def send_subject(resource, subject)
    return SUBJECT_HEAD + subject if resource.eql?('article')
    SUBJECT_HEAD + REPLY_HEAD + subject
  end
end
