module HttpClient
  class CommonResposne
    def self.success
      { statusCode: 200, message: 'success' }
    end

    def self.error
      { statusCode: 500, message: 'error' }
    end
  end
end
