import { User } from '../model/user-schema.mjs';
import { createUser } from '../handler/user-handler.mjs';
// Mock the User model
jest.mock('../model/user-schema.mjs'); // Adjust the path to your User model

// Mock express request and response objects
const mockRequest = body => ({
  body,
  // If you use `express-validator` for `matchedData`, mock it here
  params: {},
  query: {},
});

const mockResponse = () => {
  const res = {};
  res.status = jest.fn().mockReturnValue(res);
  res.json = jest.fn().mockReturnValue(res);
  return res;
};

describe('createUser', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('should return 400 if user already exists', async () => {
    // Arrange
    const existingUser = { username: 'testuser' };
    User.findOne.mockResolvedValue(existingUser);

    const req = mockRequest({
      username: 'testuser',
      displayName: 'Test User',
      password: 'password123',
    });
    const res = mockResponse();

    // Act
    await createUser(req, res);

    // Assert
    expect(res.status).toHaveBeenCalledWith(400);
    expect(res.json).toHaveBeenCalledWith({ message: 'User already exists' });
  });

  it('should return 201 and create a new user if user does not exist', async () => {
    // Arrange
    User.findOne.mockResolvedValue(null); // Simulate user not existing
    User.prototype.save = jest.fn().mockResolvedValue({
      username: 'newuser',
      displayName: 'New User',
      password: 'hashedpassword',
    });

    const req = mockRequest({
      username: 'newuser',
      displayName: 'New User',
      password: 'password123',
    });
    const res = mockResponse();

    // Act
    await createUser(req, res);

    // Assert
    expect(res.status).toHaveBeenCalledWith(201);
    expect(User).toHaveBeenCalledWith({
      username: 'newuser',
      displayName: 'New User',
      password: 'hashedpassword',
    });

    // expect(res.json).toHaveBeenCalledWith({
    //   user: { username: 'newuser', displayName: 'New User' },
    // });
  });

  // it('should return 500 on error', async () => {
  //   // Arrange
  //   User.findOne.mockRejectedValue(new Error('Database error'));

  //   const req = mockRequest({
  //     username: 'erroruser',
  //     displayName: 'Error User',
  //     password: 'password123',
  //   });
  //   const res = mockResponse();

  //   // Act
  //   await createUser(req, res);

  //   // Assert
  //   expect(res.status).toHaveBeenCalledWith(500);
  //   expect(res.json).toHaveBeenCalledWith({ error: 'Error: Database error' });
  // });
});
