class ArticlesController < ApplicationController
  before_action :validate_article_id, only: [:show, :destroy]

  def index
    @articles = Article.only_visible
  end

  def show
    @article = Article.find(path_params[:id])
  end

  def new
    @article = Article.new
  end

  def create
    if Article.create_record(article_params)
      redirect_to action: :index
    else
      redirect_to action: :new
    end
  end

  def destroy
    Article.archive(path_params[:id])
    redirect_to action: :index
  end

  private

  def validate_article_id
    id = path_params[:id]
    raise if id.blank? || !Article.exists?(id)
  end

  def path_params
    params.permit(:id)
  end

  def article_params
    params.require(:article).permit(:id, :title, :text)
  end
end
