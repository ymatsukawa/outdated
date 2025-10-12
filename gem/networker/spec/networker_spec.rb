# frozen_string_literal: true

require_relative "../lib/networker/http"
RSpec.describe Networker do
  describe "#gem_version" do
    it "has a version number" do
      expect(Networker.gem_version).not_to be nil
    end
  end

  describe "Networker::Http#req" do
    context "when request by file" do
      subject { Networker::Http.req(opt_key, opt_val) }
      let(:opt_key) { :file }
      where(:file_name) do
        [
          "spec/FILES/http_get.json",
          "spec/FILES/http_get_query.json",
          "spec/FILES/http_post.json",
          "spec/FILES/http_put.json",
        ]
      end
      with_them do
        let(:opt_val) { file_name }

        it "should run process" do
          expect { subject }.not_to raise_error
        end
      end

    end
  end
end
