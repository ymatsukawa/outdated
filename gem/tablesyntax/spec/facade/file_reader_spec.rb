# frozen_string_literal: true

require "pathname"
require_relative "../../lib/tablesyntax/facade/file_reader"
require_relative "../../lib/tablesyntax/fileio/output/standard_output"

RSpec.describe Tablesyntax::Facade::FileReader do
  shared_examples_for "should read file as yaml" do
    let(:path) { Pathname.new("spec/FILES/basic.yml").realpath.to_s }
    let(:read_stream) { File.read(path) }
    before do
      allow(File).to receive(:read).with(path).and_return(read_stream)
      allow(Psych).to receive(:safe_load).and_return(loaded_map)
      allow(Tablesyntax::Fileio::Output::StandardOutput).to receive(:out)
    end

    it do
      subject

      expect(File).to have_received(:read).with(path)
      expect(Psych).to have_received(:safe_load)
      expect(Tablesyntax::Fileio::Output::StandardOutput).to have_received(:out)
    end
  end

  shared_examples_for "should read file as json" do
    let(:path) { Pathname.new("spec/FILES/basic.json").realpath.to_s }
    let(:read_stream) { File.read(path) }
    before do
      allow(File).to receive(:read).with(path).and_return(read_stream)
      allow(JSON).to receive(:parse).with(read_stream).and_return(loaded_map)
      allow(Tablesyntax::Fileio::Output::StandardOutput).to receive(:out)
    end

    it do
      subject

      expect(File).to have_received(:read).with(path)
      expect(JSON).to have_received(:parse).with(read_stream)
      expect(Tablesyntax::Fileio::Output::StandardOutput).to have_received(:out)
    end
  end

  describe "#run" do
    before do
      allow(Tablesyntax::Facade::FileReader).to receive(:open)
    end
    subject { Tablesyntax::Facade::FileReader.run(path) }

    context "when path is valid" do
      let(:loaded_map) do
        {
          "keys" => %i[title content],
          "table_kvs" => [
            {
              title: "title_example",
              content: "content_example"
            }
          ]
        }
      end
      it_behaves_like "should read file as yaml"
      it_behaves_like "should read file as json"
    end

    context "when path is invalid" do
      where(:path_) do
        [
          "spec/FILES/unspported.dat",
          "/tmp/im.not.exist.file"
        ]
      end
      with_them do
        let(:path) { Pathname.new(path_).realpath.to_s }

        it "should raise error" do
          expect { subject }.to raise_error StandardError
        end
      end
    end

  end
end
