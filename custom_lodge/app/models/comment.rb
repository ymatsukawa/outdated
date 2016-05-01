class Comment < ActiveRecord::Base
  belongs_to :user
  belongs_to :article
  counter_culture :article
  after_create :create_notification_by_create
  before_update :create_notification_by_update

  validates :body, presence: true

  # added @ custom_v0.0.1
  def reply_title
    Settings.lodge.comment.reply_title.head + self.article.title
  end

  def create_notification(state)
    article = self.article
    notification = CommentNotification.create!(
      user_id: self.user_id,
      state: state,
      article_id: article.id,
    )
    notification.create_targets_for_stocked_user_by_article(article)
  end

  def create_notification_by_create
    create_notification(:create)
  end

  def create_notification_by_update
    create_notification(:update)
  end

  def remove_user_notification(current_user)
    notifications = ArticleNotification.where(article_id: self.id)
    NotificationTarget.destroy_all(notification_id: notifications.map {|n| n.id }, user_id: current_user.id)
    notifications.each do |notification|
      notification.destroy! if notification.notification_targets.length == 0
    end
  end
end
