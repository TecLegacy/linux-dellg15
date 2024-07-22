import { getUserById } from '../handler/user-handler.mjs';
import { User } from '../model/user-schema.mjs';

jest.mock('../model/user-schema.mjs');

describe('getUserById', () => {
  let mockRequest;
  let mockResponse;

  beforeEach(() => {
    mockRequest = { userId: '669a83b10e4523cff1332ea4' };
    mockResponse = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    };
    jest.clearAllMocks();
  });

  it('returns 404 status with json error string when user is not found', async () => {
    jest
      .spyOn(User, 'find')
      .mockImplementationOnce(() => Promise.resolve(null));

    await getUserById(mockRequest, mockResponse);

    expect(User.find).toHaveBeenCalledWith({ _id: mockRequest.userId });
    expect(mockResponse.status).toHaveBeenCalledWith(404);
    expect(mockResponse.json).toHaveBeenCalledWith({
      error: [{ message: 'User not found' }],
    });
  });

  it('returns 500 when mongodb errors out', async () => {
    jest
      .spyOn(User, 'find')
      .mockImplementationOnce(() => Promise.reject('MongoDB error'));

    await getUserById(mockRequest, mockResponse);

    expect(User.find).toHaveBeenCalledWith({ _id: mockRequest.userId });
    expect(mockResponse.status).toHaveBeenCalledWith(500);
    expect(mockResponse.json).toHaveBeenCalledWith({
      error: [{ message: 'Internal Server error' }],
    });
  });

  it("returns 200 status with user's data", async () => {
    const mockUser = {
      _id: '669a83b10e4523cff1332ea4',
      username: 'patch',
      password: 'username@12A',
      displayName: 'awesomeGodX',
      createdAt: '2024-07-19T15:18:09.289Z',
      updatedAt: '2024-07-19T16:09:04.775Z',
      __v: 0,
    };

    jest
      .spyOn(User, 'find')
      .mockImplementationOnce(() => Promise.resolve([mockUser]));

    await getUserById(mockRequest, mockResponse);

    expect(User.find).toHaveBeenCalledWith({ _id: mockRequest.userId });
    expect(mockResponse.status).toHaveBeenCalledWith(200);
    expect(mockResponse.json).toHaveBeenCalledWith({ user: [mockUser] });
  });
});
