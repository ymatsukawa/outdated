require_relative "../../lib/networker/parser/map_file"

RSpec.describe Networker::Parser::MapFile do
  describe "#read" do
    subject { Networker::Parser::MapFile.read(path) }

    context "when path is valid" do
      where(:path_, :expected_) do
        [
          [
            "spec/FILES/http_get.json",
            { "method" => "GET", "scheme" => "http", "host" => "localhost:3000", "path" => "/example/html", "headers" => {} }
          ]
        ]
      end
      with_them do
        let(:path) { path_ }
        let(:expected) { expected_ }
        it "should read file" do
          expect(subject).to eq(expected)
        end
      end
    end

  end
end