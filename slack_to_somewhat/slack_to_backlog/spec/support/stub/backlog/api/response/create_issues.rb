require 'json'

module Stub
  module Backlog
    module Api
      module Response
        CREATE_ISSUES = {
          id: 1,
          projectId: 1,
          issueKey: 'BLG-1',
          keyId: 1,
          issueType: {
            id: 2,
            projectId: 1,
            name: 'タスク',
            color: '#7ea800',
            displayOrder: 0
          },
          summary: 'first issue',
          description: '',
          resolutions: nil,
          priority: {
            id: 3,
            name: '中'
          },
          status: {
            id: 1,
            name: '未対応'
          },
          assignee: {
            id: 2,
            userId: 'ymatsukawa',
            name: 'ymatsukawa',
            roleType: 2,
            lang: nil,
            mailAddress: 'example@example.com'
          },
          category: [],
          versions: [],
          milestone: [
            {
              id: 30,
              projectId: 1,
              name: 'wait for release',
              description: '',
              startDate: nil,
              releaseDueDate: nil,
              archived: false,
              displayOrder: 0
            }
          ],
          startDate: nil,
          dueDate: nil,
          estimatedHours: nil,
          actualHours: nil,
          parentIssueId: nil,
          createdUser: {
            id: 1,
            userId: 'admin',
            name: 'admin',
            roleType: 1,
            lang: 'ja',
            mailAddress: 'example@example.com'
          },
          created: '2012-07-23T06:10:15Z',
          updatedUser: {
            id: 1,
            userId: 'admin',
            name: 'admin',
            roleType: 1,
            lang: 'ja',
            mailAddress: 'exmaple@example.com'
          },
          updated: '2012-07-23T06:10:15Z',
          customFields: [],
          attachments: [
            {
              id: 1,
              name: 'IMGP0088.JPG',
              size: 85079
            },
          ],
          sharedFiles: [],
          stars: []
        }.to_json
      end
    end
  end
end
