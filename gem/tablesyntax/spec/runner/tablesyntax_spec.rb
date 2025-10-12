# frozen_string_literal: true

RSpec.describe Tablesyntax do
  describe "#gem_version" do
    it "has a version number" do
      expect(Tablesyntax.gem_version).not_to be nil
    end
  end
end
