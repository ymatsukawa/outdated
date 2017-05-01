import { requestCreateTicketToBacklog }   from '../../actions'
import { successMessages, errorMessages } from '../mocks/requestMessageIF'

describe('requestCreateTicketToBacklog(message)', () => {
  let errorMessage
  beforeAll(() => {
    errorMessage = 'your format may be illegal. check official docs.'
  })

  describe('when message is valid', () => {
    for(let [testCase, requestMessage] of Object.entries(successMessages)) {
      test(testCase, () => {
        const response = requestCreateTicketToBacklog(requestMessage)
        expect(response.statusCode).toBe(200)
        expect(response.message).toBe('http://localhost')
      })
    }
  })

  describe('when message is invalid', () => {
    for(let [testCase, requestMessage] of Object.entries(errorMessages)) {
      test(testCase, () => {
        const response = requestCreateTicketToBacklog(requestMessage)
        expect(response.statusCode).toBe(500)
        expect(response.message).toBe(errorMessage)
      })
    }
  })
})
