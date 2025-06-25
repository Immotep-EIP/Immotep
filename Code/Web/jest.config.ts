export default {
  testEnvironment: 'jest-environment-jsdom',
  transform: {
    '^.+\\.tsx?$': [
      'ts-jest',
      {
        diagnostics: false
      }
    ]
  },
  moduleNameMapper: {
    '\\.(css|less|sass|scss)$': 'identity-obj-proxy',
    '\\.(jpg|jpeg|png|gif|webp|svg)$': '<rootDir>/__mocks__/fileMock.js',
    '^@/(.*)$': '<rootDir>/src/$1'
  },
  setupFilesAfterEnv: ['<rootDir>/setupTests.ts'],
  testPathIgnorePatterns: [
    '<rootDir>/src/enums/',
    '<rootDir>/src/utils/types/'
  ],
  coveragePathIgnorePatterns: [
    '<rootDir>/src/enums/',
    '<rootDir>/src/utils/types/'
  ]
}
