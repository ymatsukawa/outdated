class V1::ArticlesController < ApplicationController
  def create
    create_article = V1::CreateArticle.new(create_params)

    render json: create_article.as_json
  end

  def create_params
    params.require(:article).permit(:title, :text)
  end
end