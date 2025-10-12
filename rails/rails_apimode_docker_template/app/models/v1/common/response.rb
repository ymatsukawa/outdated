class V1::Common::Response
  STATUS = {
    success: 200, created: 201,  bad_request: 400, unauthorized: 401, server_error: 500
  }
  
  def self.create_json(status, body = nil)
    message =
      case(status)
        when STATUS[:success]
          'success'
        when STATUS[:created]
          'resource created'
        when STATUS[:bad_request]
          'bad request'
        when STATUS[:unauthorized]
          'unauthorized'
        when STATUS[:server_error]
          'internal server error'
      end
    
    json = {
      status: status,
      message: message
    }

    json.merge!(body: body) if body.present?
    json
  end

  def self.create_error_json(errors, status_code = 400)
    create_json(status_code, { 'errors': errors })
  end
end