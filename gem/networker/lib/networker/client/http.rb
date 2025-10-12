# frozen_string_literal: true

require "faraday"
require "faraday/net_http"

module Networker
  module Client
    class Http
      def initialize(url, req_headers, params)
        Faraday.default_adapter = :net_http

        @url = url
        @req_headers = req_headers
        @params = params

        @status = -1
        @headers = ""
        @body = ""
      end

      def get
        self.request(:get)
      end

      def post
        self.request_with_body(:post)
      end

      def put
        self.request_with_body(:put)
      end

      def delete
        self.request_with_body(:delete) 
      end

      attr_reader :status, :headers, :body

      private

      def response=(res)
        @status = res.status
        @headers = res.headers
        @body = res.body
      end

      def request(method)
        self.response = Faraday.new(url: @url, params: @params, headers: @req_headers).send(method)
      end

      def request_with_body(method)
        self.response = Faraday.new(url: @url, headers: @req_headers).send(method) do |f|
          f.body = @params unless @params.nil?
        end
      end

      attr_reader :url, :req_headers, :params
    end
  end
end
