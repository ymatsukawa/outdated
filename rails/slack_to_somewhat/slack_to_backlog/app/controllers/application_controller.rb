class ApplicationController < ActionController::API
  rescue_from Exception,                      with: :render_error500
  rescue_from ActionController::RoutingError, with: :render_error404

  def routing_error
    raise ActionController::RoutingError, params[:path]
  end

  private
  
  def render_error404(e)
    logger.info e.message if e.present?
    render json: { status_code: 404, message: 'resource not found' }, status: :not_found
  end
  
  def render_error500(e)
    logger.info e.message if e.present?
    render json: { status_code: 500, message: 'server error. contact server administrator.' }, status: :internal_server_error
  end
end
