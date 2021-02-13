require 'backlog_kit'

module Backlog
  class Base
    include Response

    def initialize
      @backlog_api_client = BacklogKit::Client.new(
        space_id: C.api.backlog.space_id,
        api_key:  C.api.backlog.api_key
      )
    end
  end
end
