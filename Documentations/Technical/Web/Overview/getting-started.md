# Getting Started with Keyz

This guide will help you set up and run the Keyz web application on your local machine.

## Prerequisites

Before you begin, ensure you have the following installed:
- Node.js (v16.x or higher)
- npm (v8.x or higher) or yarn (v1.22.x or higher)
- Git

## Installation

1. **Clone the repository**
   ```bash
   git clone Keyz-EIP
   cd Keyz/Code/Web
   ```

2. **Install dependencies**
   ```bash
   npm install
   # or
   yarn install
   ```

3. **Environment Setup**
   Create a `.env` file in the root of the web directory with the following variables:
   ```env
   VITE_API_URL=your_api_url
   ```

   Create another `.env` file in the root of the server directory with the following variables:
   ```env
   PORT='3001'
   PUBLIC_URL='http://localhost:3001'
   WEB_PUBLIC_URL='http://localhost:4242'
   DATABASE_URL='postgresql://user:password@localhost:5432/mydatabase'
   SECRET_KEY='your_secret_key_here'
   OPENAI_API_KEY='your_openai_api_key_here'
   BREVO_API_KEY='your_brevo_api_key_here'
   SHOWCASE_PUBLIC_URL="http://localhost:5000"
   ```

## Development

1. **Start the development server**
   ```bash
   npm run dev
   # or
   yarn dev
   ```

> Once the development server is running, run the web application with `npm run dev`, it will be accessible at [http://localhost:4242](http://localhost:4242).


2. **Build for production**
   ```bash
   npm run build
   # or
   yarn build
   ```

## Project Structure

```
src/
├── assets/          # Static resources (images, fonts)
├── components/      # Reusable components
├── context/         # React Context providers
├── enums/          # TypeScript enumerations
├── hooks/          # Custom React hooks
├── interfaces/     # TypeScript interfaces
├── services/       # API services
├── translation/    # i18n translations
├── utils/          # Utility functions
└── views/          # Page components
```

## Key Features

- **Authentication**: User login and registration
- **Real Property Management**: Property listing and details
<!-- - **Messaging System**: Communication between users -->
- **Responsive Design**: Mobile-first approach

## Development Guidelines

1. **Code Style**
   - Follow TypeScript best practices
   - Use functional components with hooks
   - Implement proper error handling
   - Write meaningful comments
   - Adhere to Airbnb JavaScript/TypeScript Style Guide
     - ESLint configuration based on `eslint-config-airbnb`
     - Prettier for consistent code formatting
     - Import ordering and naming conventions
     - Component and file structure guidelines

2. **Git Workflow**
   - Create feature branches from `dev`
   - Follow conventional commits
   - Submit pull requests for review

3. **Testing**
   ```bash
   npm run test
   # or
   yarn test
   ```

## Common Issues and Solutions

1. **Port already in use**
   - Kill the process using the port
   - Or use a different port: `npm run dev -- --port 3000`

2. **Build errors**
   - Clear node_modules: `rm -rf node_modules`
   - Reinstall dependencies: `npm install`

## Additional Resources

- [React Documentation](https://reactjs.org/docs/getting-started.html)
- [TypeScript Documentation](https://www.typescriptlang.org/docs/)
- [Vite Documentation](https://vitejs.dev/guide/)

## Support

For technical support or questions:
- Create an issue in the repository
- Contact the development team
