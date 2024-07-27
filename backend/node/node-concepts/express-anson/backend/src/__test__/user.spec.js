import { createUser } from '../handler/user-handler.mjs';
import { User } from '../model/user-schema.mjs';

// Mock MongoDB interaction
jest.mock('../model/user-schema.mjs');

jest.mock('express-validator', () => ({
  matchedData: jest.fn(() => ({
    username: 'tec',
    displayName: 'awesomeTec',
    password: 'awesome@12A',
  })),
}));

let mockRequest, mockResponse;

beforeEach(() => {
  mockRequest = {
    body: {
      username: 'tec',
      displayName: 'awesomeTec',
      password: 'awesome@12A',
    },
  };

  mockResponse = {
    status: jest.fn().mockReturnThis(),
    json: jest.fn(),
  };
});

describe('createUser function', () => {
  it('returns 400 when the user already exists in db', async () => {
    // Arrange
    User.findOne.mockResolvedValueOnce({ username: 'tec' }); // Simulate existing user

    // Act
    await createUser(mockRequest, mockResponse);

    // Assert
    expect(mockResponse.status).toHaveBeenCalledWith(400);
    expect(mockResponse.json).toHaveBeenCalledWith({
      message: 'User already exists',
    });
  });

  it("returns 201 when the user doesn't exist in db", async () => {
    // Arrange
    User.findOne.mockResolvedValueOnce(null); // Simulate no existing user

    jest.spyOn(User.prototype, 'save').mockResolvedValueOnce({
      username: 'tec',
      displayName: 'awesomeTec',
      password: 'awesome@12A',
    });

    // Act
    await createUser(mockRequest, mockResponse);

    // Assert
    expect(User).toHaveBeenCalledWith({
      displayName: 'awesomeTec',
      password: 'awesome@12A',
      username: 'tec',
    });
    expect(User.prototype.save).toHaveBeenCalled();
    expect(mockResponse.status).toHaveBeenCalledWith(201);
    expect(mockResponse.json).toHaveBeenCalledWith({
      user: {
        username: 'tec',
        displayName: 'awesomeTec',
      },
    });
  });

  it('returns 500 when there is an error during user creation', async () => {
    // Arrange
    User.findOne.mockResolvedValueOnce(null); // Simulate no existing user
    User.prototype.save.mockRejectedValueOnce(new Error('Database error'));

    // Act
    await createUser(mockRequest, mockResponse);

    // Assert
    expect(mockResponse.status).toHaveBeenCalledWith(500);
    expect(mockResponse.json).toHaveBeenCalledWith({
      error: 'Error: Database error',
    });
  });
});
