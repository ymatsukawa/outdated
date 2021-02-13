module HttpClient
  module RequestService

    class CreateTicket
      attr_accessor :body
      def initialize(body)
        @body = body
      end

      def request(context)
        context.client.post do |req|
          req.url '/backlog/ticket'
          req.headers['Content-Type'] = 'application/json'
          req.body = @body
        end
      end
    end

  end
end
