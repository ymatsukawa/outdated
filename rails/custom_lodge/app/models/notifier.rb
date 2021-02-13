# added @ custom_v0.0.1
class Notifier
  class << self
    def send_notice_email(article, sub_resource = 'article')
      bodies = email_bodies(article, sub_resource) # out to avoid wasted execution
      User.notice_email_receivers.each do |receiver|
        Notify.notice_email(receiver.email, sub_resource, bodies).deliver_later!
      end
    end

    private

    def email_bodies(article, sub_resource)
      bodies = { article_id: article.id, subject: article.title }
      case(sub_resource)
      when 'article'
        bodies.merge(body: article.body)
      when 'comment'
        bodies.merge(body: article.comments.last.body)
      else
        raise LodgeErrors::ImplementationError, Settings.lodge.notifier.error.implementation_error
      end
    end
  end
end
