require_relative './message_formatter/create_ticket'

module Backlog
  class Ticket < Base
    module MESSAGE_INDEX
      PROJECT_ID = 1
      TITLE      = 2
      SUBJECT    = 3
      PRIORITY   = 4
      ISSUE_TYPE = 5
    end
    attr_accessor :message

    attr_accessor :project_id
    attr_accessor :summary
    attr_accessor :priority_id
    attr_accessor :issue_type_id


    def self.new_instance(params)
      obj = self.new
      obj.message = params[:message]
      obj
    end

    def valid
      MessageFormatter::CreateTicket::FORMAT
      if MessageFormatter::CreateTicket::FORMAT.match(@message)
        return true
      end
      false
    end

    def create
      return error unless valid
      begin
        # should respond
        # @backlog_api_client.create_issue(
        #   @summary, { project_id: @project_id, priority_id: @priority_id, issue_type_id: @issue_type_id }
        # )
        success
      rescue => e
        error
      end
    end

  end
end
