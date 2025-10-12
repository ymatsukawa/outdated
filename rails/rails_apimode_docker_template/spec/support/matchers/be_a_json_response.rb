RSpec::Matchers.define :be_a_json_response do |expected_body|
  failure = nil

  match do |actual|
    header_matcher = have_attributes(content_type: 'application/json; charset=utf-8', charset: 'utf-8')
    result = header_matcher.matches?(actual)
    failure = header_matcher.failure_message unless result

    body_matcher = if expected_body.respond_to?(:matches?)
                     expected_body
                   elsif !expected_body.nil?
                     eq(expected_body)
                   end
    if body_matcher.nil?
      result
    else
      begin
        json = with_indifferent_access(JSON.parse(actual.body))
        body_matcher.matches?(json).tap do |body_result|
          failure = [failure, body_matcher.failure_message].compact.join(', ') unless body_result
        end && result
      rescue JSON::ParserError => e
        failure = [failure, "#{e.class.name}: #{e.message}"].compact.join(', ')
        false
      rescue TypeError => e
        failure = [failure, "#{e.class.name}: #{e.message}\n#{actual.body}"].compact.join(', ')
        false
      end
    end
  end

  failure_message { failure || '' }

  def with_indifferent_access(json)
    case json
    when Hash
      json.with_indifferent_access.transform_values {|v| with_indifferent_access(v) }
    when Array
      json.map {|element| with_indifferent_access(element) }
    else
      json
    end
  end
end