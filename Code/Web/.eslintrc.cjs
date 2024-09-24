module.exports = {
    parser: '@typescript-eslint/parser',
    extends: ['airbnb', 'airbnb/hooks'],
    plugins: [],
    settings: {
      'import/resolver': {
        typescript: {},
      },
    },
    env: {
      browser: true,
      es2021: true,
    },
    rules: {
      'react/jsx-filename-extension': [1, { extensions: ['.jsx', '.tsx'] }],
      'max-len': [
        'error',
        {
          code: 120,
          ignoreComments: true,
          ignoreTemplateLiterals: true,
          ignoreRegExpLiterals: true,
          ignoreUrls: true,
        },
      ],
      'react/require-default-props': 'off',
      'import/no-unresolved': 'off',
      'import/extensions': 'off',
      'import/no-extraneous-dependencies': 'off',
      "react/react-in-jsx-scope": "off",
      'react/function-component-definition': ['error', {
        namedComponents: 'arrow-function', unnamedComponents: 'arrow-function',
      }],
      'react/jsx-props-no-spreading': 'off',
      'no-unused-vars': 'off',
      'no-shadow': 'off',
      'no-console': ['warn', { allow: ['error'] }],
      'react-hooks/exhaustive-deps': 'off',
    },
    overrides: [
      {
        files: ['*.ts', '*.tsx'],
        rules: {
          'import/no-extraneous-dependencies': 'off',
        },
      },
    ],
  };