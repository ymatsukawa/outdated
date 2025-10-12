module RequestSpecHelper
  shared_examples_for 'response with status' do |expected_status|
    it 'should response with status' do
      message = expected_message(expected_status)

      send_request
      expect(response).to be_a_json_response(including(status: expected_status))
    end
  end

  shared_examples_for 'response with body and status' do |expected_status|
    it 'should response with body' do
      message = expected_message(expected_status)

      send_request
      expect(response).to be_a_json_response(including(status: expected_status))
      expect(response).to be_a_json_response(including(message: message))
      expect(response).to be_a_json_response(including(expected_body))
    end
  end

  shared_examples_for 'response without body and status' do |expected_status|
    it 'should response without body' do
      message = expected_message(expected_status)

      send_request
      expect(response).to be_a_json_response(including(status: expected_status))
      expect(response).to be_a_json_response(including(message: message))
      expect(response).not_to be_a_json_response(including(:body))
    end
  end

  def expected_message(status_code)
    case(status_code)
      when 200
        'success'
      when 201
        'resource created'
      when 400
        'bad request'
      when 401
        'unauthorized'
      when 500
        'server error'
    end
  end
end