export const GETuserQuerySchema = {
  filter: {
    in: ['query'],
    optional: true,
    isString: {
      errorMessage: 'Filter must be a string',
    },
    notEmpty: {
      errorMessage: 'Filter cannot be empty',
    },
  },
  sort: {
    in: ['query'],
    optional: true,
    isIn: {
      options: [['displayName', 'username']],
      errorMessage: 'Sort must be displayName or username',
    },
  },
};

export const POSTuserBodySchema = {
  username: {
    isString: {
      errorMessage: 'Username must be a string',
    },
    notEmpty: {
      errorMessage: 'Username cannot be empty',
    },
  },
  displayName: {
    isString: {
      errorMessage: 'Display Name must be a string',
    },
    notEmpty: {
      errorMessage: 'Display Name cannot be empty',
    },
  },
  password: {
    isString: {
      errorMessage: 'Password must be a string',
    },
    notEmpty: {
      errorMessage: 'Password cannot be empty',
    },
    isStrongPassword: {
      minLength: 8,
      minLowercase: 1,
      minUppercase: 1,
      minNumbers: 1,
      minSymbols: 1,
      errorMessage:
        'Password must be at least 8 characters long and include at least one uppercase letter, one lowercase letter, one number, and one symbol',
    },
  },
};

export const PUTuserBodySchema = {
  username: {
    isString: {
      errorMessage: 'Username must be a string',
    },
    optional: true,
  },
  displayName: {
    isString: {
      errorMessage: 'Display Name must be a string',
    },
    optional: true,
  },
  password: {
    isString: {
      errorMessage: 'Password must be a string',
    },
    optional: true,
    isStrongPassword: {
      minLength: 8,
      minLowercase: 1,
      minUppercase: 1,
      minNumbers: 1,
      minSymbols: 1,
      errorMessage:
        'Password must be at least 8 characters long and include at least one uppercase letter, one lowercase letter, one number, and one symbol',
    },
  },
};
