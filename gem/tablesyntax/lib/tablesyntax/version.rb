# frozen_string_literal: true

module Tablesyntax
  def self.gem_version
    VERSION::CODE
  end

  module VERSION
    MAJOR = 0
    MINOR = 0
    TINY  = 3

    CODE = [MAJOR, MINOR, TINY].compact.join(".")
  end
end
