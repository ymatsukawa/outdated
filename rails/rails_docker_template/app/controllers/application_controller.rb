class ApplicationController < ActionController::Base
  rescue_from RuntimeError, with: :render_500

  private

  def render_500
    render 'errors/internal_server_error', status: 500
  end
end
