export const userQuerySchema = {
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

export const userBodySchema = {
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
};
