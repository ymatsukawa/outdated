class V1::CreateArticle
  include ActiveModel::Validations
  validates :title, presence: true, length: { maximum: 255 }
  validates :text, presence: true, length: { maximum: 1000 }

  def initialize(attrs)
    @title = attrs[:title]
    @text = attrs[:text]
  end

  def as_json
    if !valid?
      return V1::Common::Response.create_error_json(@errors.full_messages, 400)
    end

    record = {
      title: @title,
      text: @text
    }

    begin
      Article.create!(record)
      body = response_body(record)
      return V1::Common::Response.create_json(201, body)
    rescue => e
      return V1::Common::Response.create_json(500)
    end
  end

  private

  def response_body(attrs)
    {
      article: {
        title: attrs[:title]
      }
    }
  end

  attr_reader :title
  attr_reader :text
end