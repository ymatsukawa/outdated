module MessageFormatter
  module CreateTicket
    # do not use heredocs because last newline will be contained.
    FORMAT = Regexp.new(
'create ticket__;
プロジェクト番号\[\d+\]__;
タイトル\[[\w\W]+\]__;
概要\[[\w\W]+\]__;
優先度\[.+\]__;
種別\[.+]__;'
)
  end
end
