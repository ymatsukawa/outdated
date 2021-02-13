require 'json'

module Stub
  module Backlog
    module Api
      module Common
        END_POINT = C.api.backlog.endpoint
        API_KEY   = C.api.backlog.api_key
        REQUEST_HEADER = {
          'Accept'          => '*/*',
          'Accept-Encoding' => 'gzip;q=1.0,deflate;q=0.6,identity;q=0.3',
          'Content-Type'    => 'application/x-www-form-urlencoded',
          'User-Agent'      => 'BacklogKit Ruby Gem 0.14.0'
        }

        module Response
          ERROR = {
            errors: [
              {
                message: 'invalid request to backlog api',
                code: 6,
                moreInfo: ''
              }
            ]
          }.to_json
        end
      end
    end
  end
end
