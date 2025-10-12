module RequestSpecHelper
  shared_examples_for 'response with status code' do |expected_status|
    it 'should response with status' do
      send_request
      expect(response).to have_http_status(expected_status)
    end
  end
end