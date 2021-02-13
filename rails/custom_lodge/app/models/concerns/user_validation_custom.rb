# added custom concern @ custom_v0.0.1
module UserValidationCustom
  extend ActiveSupport::Concern

  included do
    [
      { only_integer: true },
      { greater_than_or_equal_to: 0 },
      { less_than_or_equal_to:    1 },
    ].each do |validate_detail|
      validates_numericality_of :notice_mail_send_state, validate_detail
    end
  end
end
