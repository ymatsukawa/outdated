require_relative "../../lib/networker/client/http"

RSpec.describe Networker::Client::Http do
  describe "#get, #post, #put, #delete" do
    let(:client) { Networker::Client::Http.new(url, req_headers, params) }

    context "when request is valid" do
      where(:method, :url_, :params_, :expected_body_message, :expected_status) do
        [
          [:get,    "http://localhost:3000/example/html",    nil,                                 "This is for HTML Test", 200],
          [:post,   "http://localhost:3000/example/machine", { nut: 3, iron_plate: 5 },           "machine created",       201],
          [:put,    "http://localhost:3000/example/clothes", { necktie: "gray", suit: "stripe" }, "clothes updated",       200],
          [:delete, "http://localhost:3000/example/paper",   { paper_id: 123 },                   "paper deleted",         200]
        ]
      end
      with_them do
        subject { client.send(method) }
        let(:url) { url_ }
        let(:req_headers) { {} }
        let(:params) { params_ }
    
        it "should get success response html" do
          subject
          expect(client.status).to eq expected_status
          expect(client.headers).to include({ "content-type" => "text/html;charset=utf-8" })
          expect(client.body).to include expected_body_message
        end
      end

      context "when queries are given" do
        subject { client.get }
        let(:url) { "http://localhost:3000/example/html" }
        let(:req_headers) { {} }
        let(:params) { { first_name: "yamada", last_name: "taro", additional: "more-query" } }
  
        it "should response with passed query" do
          subject
          ["yamada", "taro"].each do |expected|
            expect(client.body).to include expected
          end
          expect(client.body).not_to include "more-query"
        end
      end
    end

    context "when request is invalid" do
      where(:method, :url_, :params_, :expected_status) do
        [
          [:get,    "http://localhost:3000/blahblah",        nil,                     404],
          [:post,   "http://localhost:3000/example/machine", { tuna: 3 },             400],
          [:put,    "http://localhost:3000/example/clothes", { swim_suit: "stripe" }, 400],
          [:delete, "http://localhost:3000/example/paper",   { fever_id: 123 },       400]
        ]
      end
      with_them do
        subject { client.send(method) }
        let(:url) { url_ }
        let(:req_headers) { {} }
        let(:params) { params_ }
    
        it "should get fail response" do
          subject
          expect(client.status).to eq expected_status
          expect(client.headers).to include({ "content-type" => "text/html;charset=utf-8" })
        end
      end
    end
  end

end