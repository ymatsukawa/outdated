require_relative "../../lib/networker/facade/http"
require_relative "../../lib/networker/client/http"

RSpec.describe Networker::Facade::Http do
  describe "#req" do
    let(:client) { double(Networker::Client::Http) }
    subject { Networker::Facade::Http.req(options) }

    context "when options are valid" do
      let(:url) { [options["scheme"], options["host"]].join("://") + options["path"] }
      let(:headers) { options["headers"] }
      let(:req_options) { options["params"] }

      where(:method, :mock_status) do
        [
          [:get,    200],
          [:post,   201],
          [:put,    200],
          [:delete, 200]
        ]
      end
      with_them do
        let(:options) do
          {
            "method" => method.to_s.upcase,
            "scheme" => "http",
            "host" => "example.com",
            "path" => "/example",
            "headers" => {}
          }
        end
        let(:response) do
          {
            status: mock_status,
            headers: "Content-Type: text/html; charset=utf-8",
            body: "<html><body><h1>mock response</h1></body></html>"
          }
        end
        before do
          allow(Networker::Client::Http).to receive(:new).with(url, headers, req_options).and_return(client)
          allow(client).to receive(method)
          [:status, :headers, :body].each do |mock_key|
            allow(client).to receive(mock_key).and_return(response[mock_key])
          end
        end
        
        it do
          subject

          expect(Networker::Client::Http).to have_received(:new).with(url, headers, req_options)
          expect(client).to have_received(method)
        end
      end
    end

    context "when options are invalid" do
      context "when method does not support" do
        let(:options) do
          { method: "INVALID" }
        end
        
        it "should raise error" do
          expect { subject }.to raise_error
        end
      end
    end
  end
end