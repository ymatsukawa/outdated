module HttpClient
  module RequestService

    class ProjectInfo
      attr_accessor :body
      def initialize(body)
        @body = body
      end

      def request(context)
        context.client.get do |req|
          req.url '/backlog/info'
          req.headers['Content-Type'] = 'application/json'
          req.body = @body
        end
      end
    end

  end
end
