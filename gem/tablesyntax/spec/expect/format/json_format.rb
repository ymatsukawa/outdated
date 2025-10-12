# frozen_string_literal: true

module Expect
  module Format
    class JsonFormat
      BASIC_FORMAT = <<~EXPECT
where(:title, :content) do
 "json_test_title_0" |     "json_content_0"
"this is title test" | "this is content test"
end
EXPECT

      MANY_TYPES_FORMAT = <<~EXPECT
where(:title, :metadata) do
        "string" |    "i am string"
       "integer" |                1
         "float" |             3.14
         "array" | [1, 2, "a", "b"]
"array in array" | [1, 2, ["a", "b"]]
           "map" | {"a"=>1, "b"=>2}
    "map in map" |  {"a"=>{"b"=>2}}
 "repeat string" |   variable * 987
      "show nil" |              nil
end
EXPECT
    end
  end
end