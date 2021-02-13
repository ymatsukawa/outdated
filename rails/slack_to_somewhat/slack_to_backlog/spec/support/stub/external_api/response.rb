module Stub
  module ExternalApi
    module Response
      SUCCESS         = { status_code: 200, message: 'success' }
      ERROR           = { status_code: 500, message: 'internal server error' }
      NOT_FOUND_ERROR = { status_code: 404, message: 'resource not found' }
    end
  end
end
