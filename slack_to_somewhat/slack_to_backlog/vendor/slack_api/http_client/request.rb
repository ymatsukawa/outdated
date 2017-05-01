module HttpClient
  class Request
    attr_accessor :client
    attr_accessor :request_scheme
    def initialize(request_scheme)
      @client =
        Faraday.new('http://localhost:3000') do |faraday|
          faraday.request :json
          faraday.adapter :net_http
        end
      @request_scheme = request_scheme
    end

    def request
      @request_scheme.request(self)
    end
  end
end
