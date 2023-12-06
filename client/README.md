# TinyURL Client

## Project setup
1. Copy over `.env.default`
    ```
    cp .env.default .env
    ```
2. Update the `VITE_APP_SERVER_URL` variable to be the same as the URL of your server. By default it is set to the
   port on which the server will listen to if started using Docker.
3. Install dependencies
    ```
    # yarn
    yarn

    # npm
    npm install

    # pnpm
    pnpm install

    # bun
    bun install
    ```
4. Start front-end
    ```
    # yarn
    yarn dev

    # npm
    npm run dev

    # pnpm
    pnpm dev

    # bun
    bun run dev
    ```
