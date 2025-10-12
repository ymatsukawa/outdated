# frozen_string_literal: true

module Tablesyntax
  module Fileio
    module Output
      class StandardOutput
        def self.out(obj)
          puts obj
        end
      end
    end
  end
end
