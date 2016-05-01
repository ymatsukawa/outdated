class NoticeDeliverJob < ActiveJob::Base
  queue_as :default

  def perform(email, post_obj, bodies)
    Notify.notice_email(email, post_obj, bodies).deliver_now
  end
end
