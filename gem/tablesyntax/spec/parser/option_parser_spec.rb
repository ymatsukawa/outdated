# frozen_string_literal: true

require_relative "../../lib/tablesyntax/parser/option_parser"

RSpec.describe Tablesyntax::Parser::OptionParser do
  describe "#parse and #valid" do
    context "when options are valid" do
      where(:opts) do
        [
          [[:file, "./spec/FILES/basic.yml"]],
          [[:file, 12345]]
        ]
      end
      with_them do
        let(:parser) { Tablesyntax::Parser::OptionParser.new(opts) }

        it "should parse and set attributes" do
          parser.parse

          expect(parser.valid?).to be_truthy
          expect(parser).to have_attributes(key: opts[0], val: opts[1])
        end
      end
    end

    context "when options are invalid" do
      where(:opts) do
        [
          [[:file]],
          [["-F"]],
          [["--oo--"]],
          [["abcdefg"]],
          [[12345]],
          [[nil]],
          [[]]

        ]
      end
      with_them do
        let(:parser) { Tablesyntax::Parser::OptionParser.new(opts) }

        it "should parse and set attributes" do
          parser.parse

          expect(parser.valid?).to be_falsey
          expect(parser).to have_attributes(key: nil, val: nil)
        end
      end
    end
  end
end
