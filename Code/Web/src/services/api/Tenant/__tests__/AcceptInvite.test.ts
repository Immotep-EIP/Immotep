import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'
import { AcceptInvite } from '../AcceptInvite'

jest.mock('@/services/api/apiCaller')

const mockedCallApi = callApi as jest.MockedFunction<typeof callApi>

describe('AcceptInvite', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should call callApi with the correct parameters', async () => {
    const mockLeaseId = '123'
    mockedCallApi.mockResolvedValueOnce(undefined)

    await AcceptInvite(mockLeaseId)

    expect(mockedCallApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: endpoints.tenant.invite.accept(mockLeaseId)
    })
  })

  it('should handle errors during invitation acceptance', async () => {
    const mockError = new Error('API call failed')
    mockedCallApi.mockRejectedValueOnce(mockError)

    const mockLeaseId = '123'
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    await expect(AcceptInvite(mockLeaseId)).rejects.toThrow('API call failed')

    expect(consoleErrorSpy).toHaveBeenCalledWith(
      'Error accepting invitation:',
      mockError
    )

    consoleErrorSpy.mockRestore()
  })
})
