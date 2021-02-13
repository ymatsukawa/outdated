require_relative '../common_response'
module HttpClient
  module RequestService

    class Error
      # do not need to use context.client
      def request(context)
        HttpClient::CommonResposne.error.merge(
          message: 'your order may be invalid. see docs http://localhost'
        )
      end
    end

  end
end
