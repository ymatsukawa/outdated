module UserClient
  module BadRequest
    class ContentError    < StandardError ; end
    class IdentifierError < StandardError ; end
    class TypeError       < StandardError ; end
    class PropertyError   < StandardError ; end
  end
end
