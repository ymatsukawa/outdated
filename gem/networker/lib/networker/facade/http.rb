# frozen_string_literal: true

require_relative "../client/http"

module Networker
  module Facade
    class Http
      def self.req(opts)
        self.request_options = opts
        @client = Networker::Client::Http.new(@url, @headers, @params)
        client_request

        { status: @client.status, headers: @client.headers, body: @client.body }
      end

      def self.client_request
        case @method
        when "GET"
          @client.get
        when "POST"
          @client.post
        when "PUT"
          @client.put
        when "DELETE"
          @client.delete
        else
          raise StandardError
        end
      end

      def self.request_options=(opts)
        @method = opts["method"]
        @url = [opts["scheme"], opts["host"]].join("://") + opts["path"]
        @headers = opts["headers"]
        @params = opts["params"]
      end

      private_class_method :client_request, :request_options=

      private

      attr_reader :client, :method, :url, :headers, :params
    end
  end
end
