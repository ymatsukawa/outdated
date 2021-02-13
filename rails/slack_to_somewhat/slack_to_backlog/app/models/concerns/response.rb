module Response
  extend ActiveSupport::Concern

  included do
    def success
      { status_code: 200, message: 'success' }
    end

    def error(message = C.api.external.response.error.internal_server_error)
      { status_code: 500, message: message }
    end
  end
end
