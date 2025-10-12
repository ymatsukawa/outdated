require_relative "../../lib/networker/http"
require_relative "../../lib/networker/parser/map_file"
require_relative "../../lib/networker/facade/http"

RSpec.describe Networker::Http do
  describe "#req" do
    subject { Networker::Http.req(opt_key, opt_val) }

    context "when options are valid" do
      context "when option means - read json file" do
        let(:opt_key) { :file }
        let(:opt_val) { "spec/FILES/http_get.json" }
        let(:request_map) do
          { "method" => "GET", "scheme" => "http", "host" => "example.com", "path" => "/" }
        end
        before do
          allow(Networker::Parser::MapFile).to receive(:read).with(opt_val).and_return(request_map)
          allow(Networker::Facade::Http).to receive(:req).with(request_map)
        end
        it "should respond as json case" do
          subject

          expect(Networker::Parser::MapFile).to have_received(:read).with(opt_val)
          expect(Networker::Facade::Http).to have_received(:req).with(request_map)
        end
      end
    end

    context "when options are invalid" do
      let(:opt_key) { :file }
      let(:opt_val) { nil }
      it "should raise error" do
        expect { subject }.to raise_error
      end
    end
  end
end