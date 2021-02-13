module NotifyHelper
  def send_uri(action, scheme, host, port, resouce_name, resouce_id)
    # prototype
    case action
    when 'show'
      scheme + '://' + host + ':' + port + '/' + resouce_name + '/' + resouce_id
    end
  end
end
