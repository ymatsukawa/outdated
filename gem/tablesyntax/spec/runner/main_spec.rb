# frozen_string_literal: true

RSpec.describe Tablesyntax::Main do
  describe "#run" do
    context "when option exists" do
      let(:options) { [] }
      subject { Tablesyntax::Main.run(:file, "./spec/FILES/basic.yml") }

      it do
        expect { subject }.not_to raise_error
      end
    end

    context "when option does not exist" do
      it do
        expect { Tablesyntax::Main.run }.to raise_error(StandardError)
      end
    end
  end
end
