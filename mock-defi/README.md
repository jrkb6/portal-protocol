# Registry-UI: Svelte Vite De-Fi Demo Application

This application serves as a demonstration of a decentralized finance (De-Fi) application, focusing on showcasing the login feature using the Portal system.

## Quick Start

### Prerequisites
Ensure you have the latest version of `npm` installed on your system.
Make sure portal system is running (ganache-cli, identity-server, simulations.sh, ipfs client etc.)

### Creating a Project
If you've just cloned this repository or are looking at this README in your newly created Svelte project, you've completed the first step!

Here's how you can create a new Svelte project:

```bash
npm create svelte@latest  # Create a new project in the current directory
npm create svelte@latest my-app  # Create a new project in 'my-app' directory
```
### Developing
To start your development journey, follow these steps:
1. Install dependencies: `npm install`
2. Start the development server:
   ```bash
   npm run dev  # Open the app in a new browser tab with '-- --open'
   ```
3. ```
   npm run build  # Creates a production build
   npm run preview  # Preview the production build
    ```
4. After the build, copy all files under the `build/` directory to `identity-server/build/` to serve them statically with the Go-Fiber framework (IdentityServer).
5. Before starting your application, ensure to set up the necessary environment variables in your .env file based on the .env.example:
### Environment Variables
```bash
VITE_REGISTRY_CONTRACT_ADDRESS=
VITE_IPFS_NODE=http://localhost:5001
VITE_SERVICE_URL=http://localhost:3000/api
```



