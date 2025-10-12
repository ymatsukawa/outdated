# frozen_string_literal: true

module Expect
  module Format
    class YamlFormat
      BASIC_FORMAT = <<~EXPECT
where(:title, :content) do
"test_title_0" | "test_content_0"
    "title_aa" |   "content_bb"
end
EXPECT

      MANY_TYPES_FORMAT = <<~EXPECT
where(:title, :metadata) do
             "string single quote" |                          "abcdefg"
             "string double quote" |                          "ABCDEFG"
                          "symbol" |                    :this_is_symbol
                         "integer" |                              12345
                           "array" |                   [1, 2, "a", "b"]
                  "array in array" |                 [1, 2, ["a", "b"]]
                       "deep hash" | {"root"=>{"child"=>{"data"=>123}}}
                   "repeat string" |                            'a' * 5
                       "shows nil" |                                nil
end
EXPECT
    end
  end
end