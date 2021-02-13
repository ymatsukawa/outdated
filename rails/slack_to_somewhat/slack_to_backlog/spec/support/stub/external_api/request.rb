module Stub
  module ExternalApi
    module Request
      def Request.header(content_type)
        {
          'CONTENT_TYPE' => { json: 'application/json', form_encoded: 'application/x-www-form-urlencoded' }[content_type],
          'ACCEPT'       => 'application/json'
        }
      end
    end
  end
end
